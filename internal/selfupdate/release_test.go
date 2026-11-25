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
)

func TestReleaseSelection(t *testing.T) {
	const data = `{"tag_name":"v4.0.10","assets":[{"name":"antigravity-i18n.exe","size":16},{"name":"SHA256SUMS.txt","size":87}]}`
	for _, tc := range []struct {
		current   string
		available bool
	}{{"4.0.9", true}, {"4.0.10", false}, {"4.1.0", false}} {
		r, err := readRelease(strings.NewReader(data), tc.current)
		if err != nil || r.Available != tc.available {
			t.Fatalf("version %s: %+v, %v", tc.current, r, err)
		}
	}
	for _, invalid := range []string{strings.Replace(data, "v4.0.10", "v4.0.10-beta.1", 1), strings.Replace(data, "SHA256SUMS.txt", "other", 1)} {
		if _, err := readRelease(strings.NewReader(invalid), "4.0.9"); err == nil {
			t.Fatal("accepted an ineligible release")
		}
	}
}

type transportFunc func(*http.Request) (*http.Response, error)

func (f transportFunc) RoundTrip(req *http.Request) (*http.Response, error) { return f(req) }

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
