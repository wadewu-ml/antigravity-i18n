package selfupdate

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestRateLimitDeadline(t *testing.T) {
	now := time.Unix(1700000000, 0).UTC()
	for _, tc := range []struct {
		after, remaining, reset string
		wait                    time.Duration
	}{
		{"120", "", "", 2 * time.Minute},
		{now.Add(10 * time.Minute).Format(http.TimeFormat), "", "", 10 * time.Minute},
		{"120", "0", "1700000600", 10 * time.Minute},
		{"", "0", "1700000300", 5 * time.Minute},
		{"invalid", "", "", time.Hour},
	} {
		headers := make(http.Header)
		headers.Set("Retry-After", tc.after)
		headers.Set("X-RateLimit-Remaining", tc.remaining)
		headers.Set("X-RateLimit-Reset", tc.reset)
		if got := rateLimitDeadline(headers, now); !got.Equal(now.Add(tc.wait)) {
			t.Errorf("%v: got %s, want %s", headers, got, now.Add(tc.wait))
		}
	}
}

func TestReleaseSelection(t *testing.T) {
	const data = `{"tag_name":"v4.0.10","assets":[{"name":"antigravity-i18n.exe","size":16},{"name":"SHA256SUMS.txt","size":87}]}`
	r, err := readRelease(strings.NewReader(data), "v4.0.10")
	if err != nil || !r.Available || r.size != 16 {
		t.Fatalf("release metadata: %+v, %v", r, err)
	}
	for _, invalid := range []string{strings.Replace(data, "v4.0.10", "v4.0.10-beta.1", 1), strings.Replace(data, "SHA256SUMS.txt", "other", 1)} {
		if _, err := readRelease(strings.NewReader(invalid), "v4.0.10"); err == nil {
			t.Fatal("accepted an ineligible release")
		}
	}
}

type transportFunc func(*http.Request) (*http.Response, error)

func (f transportFunc) RoundTrip(req *http.Request) (*http.Response, error) { return f(req) }

func TestCheckPublicRelease(t *testing.T) {
	const tag = "v4.0.10"
	const payload = "GUI package"
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(payload)))
	saved := client
	t.Cleanup(func() { client = saved })
	for _, tc := range []struct {
		current, page                       string
		available, limited, badSum, wantErr bool
	}{
		{current: "4.0.10", page: tag},
		{current: "4.1.0", page: tag},
		{current: "4.0.9", page: tag, available: true},
		{current: "4.0.9", page: tag, available: true, limited: true},
		{current: "4.0.9", page: tag, limited: true, badSum: true, wantErr: true},
		{current: "4.0.9", page: tag + "-beta.1", wantErr: true},
	} {
		apiCalls := 0
		client = &http.Client{Transport: transportFunc(func(req *http.Request) (*http.Response, error) {
			res := &http.Response{StatusCode: 200, Request: req, Header: make(http.Header), Body: http.NoBody}
			switch req.Method + " " + req.URL.String() {
			case "HEAD " + Repository + "/releases/latest":
				res.StatusCode = 302
				res.Header.Set("Location", Repository+"/releases/tag/"+tc.page)
			case "HEAD " + Repository + "/releases/tag/" + tc.page:
			case "GET " + tagAPI + tag:
				apiCalls++
				if tc.limited {
					res.StatusCode = 403
					res.Header.Set("X-RateLimit-Remaining", "0")
				} else {
					res.Body = io.NopCloser(strings.NewReader(fmt.Sprintf(`{"tag_name":"%s","body":"Notes","assets":[{"name":"%s","digest":"sha256:%s","size":%d}]}`, tag, exeAsset, hash, len(payload))))
				}
			case "GET " + Repository + "/releases/download/" + tag + "/" + sumsAsset:
				if tc.badSum {
					res.Body = io.NopCloser(strings.NewReader("invalid checksum"))
				} else {
					res.Body = io.NopCloser(strings.NewReader(hash + "  " + exeAsset))
				}
			case "HEAD " + Repository + "/releases/download/" + tag + "/" + exeAsset:
				res.ContentLength = int64(len(payload))
			case "GET " + Repository + "/releases/download/" + tag + "/" + exeAsset:
				res.Body = io.NopCloser(strings.NewReader(payload))
			default:
				t.Fatalf("unexpected update request: %s %s", req.Method, req.URL)
			}
			return res, nil
		})}
		r, err := Check(context.Background(), tc.current)
		if tc.wantErr {
			if err == nil {
				t.Fatalf("%+v: accepted invalid release", tc)
			}
			continue
		}
		if err != nil || r.Available != tc.available || r.Version != tag {
			t.Fatalf("%+v: got %+v, %v", tc, r, err)
		}
		if !tc.available && apiCalls != 0 {
			t.Fatal("unnecessary API call")
		}
		if tc.available {
			if !tc.limited && r.Notes != "Notes" {
				t.Fatal("missing release notes")
			}
			if err := r.download(context.Background(), filepath.Join(t.TempDir(), exeAsset), func(string, string) {}); err != nil {
				t.Fatal(err)
			}
		}
	}
}

func TestDownloadVerification(t *testing.T) {
	const payload = "GUI package"
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(payload)))
	saved := client
	t.Cleanup(func() { client = saved })
	var requests []string
	client = &http.Client{Transport: transportFunc(func(req *http.Request) (*http.Response, error) {
		if err := req.Context().Err(); err != nil {
			return nil, err
		}
		name := filepath.Base(req.URL.Path)
		requests = append(requests, name)
		body := payload
		if name == sumsAsset {
			body = hash + "  " + exeAsset + "\n"
		}
		return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(body))}, nil
	})}
	for _, tc := range []struct {
		name, digest string
		size         int64
		ok, cancel   bool
	}{
		{"release digest", "sha256:" + hash, int64(len(payload)), true, false},
		{"checksum attachment", "", int64(len(payload)), true, false},
		{"wrong digest", "sha256:" + strings.Repeat("a", 64), int64(len(payload)), false, false},
		{"truncated download", "sha256:" + hash, 99, false, false},
		{"cancel", "sha256:" + hash, int64(len(payload)), false, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			if tc.cancel {
				cancel()
			}
			r := Release{Version: "v4.1.0", size: tc.size, digest: tc.digest, hasSums: true}
			requests = nil
			path := filepath.Join(t.TempDir(), "download.exe")
			err := r.download(ctx, path, func(string, string) {})
			if (err == nil) != tc.ok {
				t.Fatalf("download: %v", err)
			}
			if tc.cancel && !errors.Is(err, context.Canceled) {
				t.Fatalf("cancel: %v", err)
			}
			if tc.ok {
				data, err := os.ReadFile(path)
				if err != nil || string(data) != payload {
					t.Fatalf("downloaded %q, %v", data, err)
				}
				expected := exeAsset
				if tc.digest == "" {
					expected = sumsAsset + "," + exeAsset
				}
				if strings.Join(requests, ",") != expected {
					t.Fatal(requests)
				}
			}
		})
	}
}
