package selfupdate

import (
	"bufio"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/mod/semver"
)

const (
	Repository = "https://github.com/wadewu-ml/regravity"
	tagAPI     = "https://api.github.com/repos/wadewu-ml/regravity/releases/tags/"
	exeAsset   = "regravity.exe"
	sumsAsset  = "SHA256SUMS.txt"
)

var client = &http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error {
	if req.URL.Scheme != "https" || len(via) >= 10 {
		return errors.New("无效的更新下载重定向")
	}
	return nil
}}

var errRateLimited = errors.New("GitHub 暂时限制了更新请求，请稍后重试，或打开发布页手动下载。")

// RateLimitError carries the time at which GitHub allows another check.
type RateLimitError struct {
	RetryAt time.Time
}

func (e *RateLimitError) Error() string { return errRateLimited.Error() }
func (e *RateLimitError) Unwrap() error { return errRateLimited }

func rateLimitDeadline(header http.Header, now time.Time) time.Time {
	var until time.Time
	if value := header.Get("Retry-After"); value != "" {
		if seconds, err := strconv.ParseUint(value, 10, 31); err == nil {
			until = now.Add(time.Duration(seconds) * time.Second)
		} else if date, err := http.ParseTime(value); err == nil {
			until = date
		}
	}
	if header.Get("X-RateLimit-Remaining") == "0" {
		if seconds, err := strconv.ParseInt(header.Get("X-RateLimit-Reset"), 10, 64); err == nil && time.Unix(seconds, 0).After(until) {
			until = time.Unix(seconds, 0)
		}
	}
	if !until.After(now) {
		return now.Add(time.Hour)
	}
	return until
}

type Release struct {
	Version   string `json:"version"`
	Notes     string `json:"notes"`
	Available bool   `json:"available"`
	RetryAt   int64  `json:"retryAt,omitempty"`
	size      int64
	digest    string
	hasSums   bool
}

func request(ctx context.Context, method, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Regravity")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		if res.StatusCode == http.StatusTooManyRequests || (res.StatusCode == http.StatusForbidden && (res.Header.Get("X-RateLimit-Remaining") == "0" || res.Header.Get("Retry-After") != "")) {
			return nil, &RateLimitError{RetryAt: rateLimitDeadline(res.Header, time.Now())}
		}
		return nil, fmt.Errorf("更新请求失败（HTTP %d）", res.StatusCode)
	}
	return res, nil
}

func get(ctx context.Context, url string) (io.ReadCloser, error) {
	res, err := request(ctx, http.MethodGet, url)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

func versionTag(version string) string { return "v" + strings.TrimPrefix(version, "v") }

func Check(ctx context.Context, current string) (*Release, error) {
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()
	// GitHub's public latest-release link resolves to the stable tag without
	// consuming an anonymous REST request.
	page, err := request(ctx, http.MethodHead, Repository+"/releases/latest")
	if err != nil {
		return nil, err
	}
	page.Body.Close()
	tag, ok := strings.CutPrefix(page.Request.URL.String(), Repository+"/releases/tag/")
	if !ok || !semver.IsValid(tag) || semver.Prerelease(tag) != "" || !semver.IsValid(versionTag(current)) {
		return nil, errors.New("无法比较当前工具与正式发布的版本号。")
	}
	if semver.Compare(tag, versionTag(current)) <= 0 {
		return &Release{Version: tag}, nil
	}
	body, err := get(ctx, tagAPI+tag)
	var limited *RateLimitError
	if errors.As(err, &limited) {
		// Official releases publish a checksum alongside the executable.
		r := &Release{Version: tag, Available: true, RetryAt: limited.RetryAt.UnixMilli(), Notes: "完整更新说明可在 GitHub 发布页查看。"}
		hash, err := r.checksum(ctx)
		if err != nil {
			return nil, errors.Join(err, limited)
		}
		asset, err := request(ctx, http.MethodHead, r.assetURL(exeAsset))
		if err != nil {
			return nil, errors.Join(err, limited)
		}
		asset.Body.Close()
		if asset.ContentLength <= 0 {
			return nil, errors.Join(errors.New("无法确认更新程序的文件大小。"), limited)
		}
		r.size, r.digest = asset.ContentLength, "sha256:"+hash
		return r, nil
	}
	if err != nil {
		return nil, err
	}
	defer body.Close()
	return readRelease(body, tag)
}

func readRelease(body io.Reader, tag string) (*Release, error) {
	var release struct {
		Tag               string `json:"tag_name"`
		Body              string
		Draft, Prerelease bool
		Assets            []struct {
			Name, Digest string
			Size         int64
		}
	}
	if err := json.NewDecoder(io.LimitReader(body, 1<<20)).Decode(&release); err != nil {
		return nil, err
	}
	if release.Draft || release.Prerelease || release.Tag != tag {
		return nil, errors.New("获取的发布信息与正式版本不符。")
	}
	r := &Release{Version: tag, Notes: release.Body, Available: true}
	for _, asset := range release.Assets {
		switch asset.Name {
		case exeAsset:
			r.size, r.digest = asset.Size, asset.Digest
		case sumsAsset:
			r.hasSums = true
		}
	}
	if r.size <= 0 || r.digest == "" && !r.hasSums {
		return nil, errors.New("该版本缺少 Windows 程序或 SHA-256 校验信息。")
	}
	return r, nil
}

func (r *Release) assetURL(name string) string {
	return Repository + "/releases/download/" + r.Version + "/" + name
}

func (r *Release) checksum(ctx context.Context) (string, error) {
	if r.digest != "" {
		if raw, ok := strings.CutPrefix(r.digest, "sha256:"); ok {
			return validHash(raw)
		}
		return "", errors.New("发布附件的摘要格式无效。")
	}
	body, err := get(ctx, r.assetURL(sumsAsset))
	if err != nil {
		return "", err
	}
	defer body.Close()
	scanner := bufio.NewScanner(io.LimitReader(body, 64<<10))
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) == 2 && strings.TrimPrefix(fields[1], "*") == exeAsset {
			return validHash(fields[0])
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", errors.New("校验文件中缺少 Windows 程序的 SHA-256。")
}

func validHash(raw string) (string, error) {
	b, err := hex.DecodeString(raw)
	if err != nil || len(b) != sha256.Size {
		return "", errors.New("无效的 SHA-256 校验值。")
	}
	return strings.ToLower(raw), nil
}

type progressWriter struct {
	size, received int64
	last           time.Time
	report         func(string, string)
}

func (p *progressWriter) Write(b []byte) (int, error) {
	p.received += int64(len(b))
	if time.Since(p.last) >= 200*time.Millisecond || p.received == p.size {
		p.report("正在下载更新", fmt.Sprintf("%.1f / %.1f MB (%.0f%%)", float64(p.received)/1e6, float64(p.size)/1e6, 100*float64(p.received)/float64(p.size)))
		p.last = time.Now()
	}
	return len(b), nil
}

func (r *Release) download(ctx context.Context, path string, report func(string, string)) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0700)
	if err != nil {
		return err
	}
	defer file.Close()
	expected, err := r.checksum(ctx)
	if err != nil {
		return err
	}
	body, err := get(ctx, r.assetURL(exeAsset))
	if err != nil {
		return err
	}
	defer body.Close()
	hash := sha256.New()
	n, err := io.Copy(io.MultiWriter(file, hash, &progressWriter{size: r.size, report: report}), io.LimitReader(body, r.size+1))
	if err != nil {
		return err
	}
	if n != r.size || hex.EncodeToString(hash.Sum(nil)) != expected {
		return errors.New("下载的文件与发布版本的大小或 SHA-256 不符，请重新下载。")
	}
	if err := file.Sync(); err != nil {
		return err
	}
	return file.Close()
}
