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
	"strings"
	"time"

	"golang.org/x/mod/semver"
)

const (
	Repository = "https://github.com/wadewu-ml/antigravity-i18n"
	latestAPI  = "https://api.github.com/repos/wadewu-ml/antigravity-i18n/releases/latest"
	exeAsset   = "antigravity-i18n.exe"
	sumsAsset  = "SHA256SUMS.txt"
)

var client = &http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error {
	if req.URL.Scheme != "https" || len(via) >= 10 {
		return errors.New("无效的更新下载重定向")
	}
	return nil
}}

type Release struct {
	Version   string `json:"version"`
	Notes     string `json:"notes"`
	Available bool   `json:"available"`
	size      int64
	digest    string
	hasSums   bool
}

func get(ctx context.Context, url string) (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "antigravity-i18n")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		if res.StatusCode == http.StatusTooManyRequests || (res.StatusCode == http.StatusForbidden && res.Header.Get("X-RateLimit-Remaining") == "0") {
			return nil, errors.New("GitHub 暂时限制了更新请求，请稍后重试，或打开发布页手动下载。")
		}
		return nil, fmt.Errorf("更新请求失败（HTTP %d）", res.StatusCode)
	}
	return res.Body, nil
}

func versionTag(version string) string { return "v" + strings.TrimPrefix(version, "v") }

func Check(ctx context.Context, current string) (*Release, error) {
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()
	body, err := get(ctx, latestAPI)
	if err != nil {
		return nil, err
	}
	defer body.Close()
	return readRelease(body, current)
}

func readRelease(body io.Reader, current string) (*Release, error) {
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
	if release.Draft || release.Prerelease || !semver.IsValid(release.Tag) || semver.Prerelease(release.Tag) != "" || !semver.IsValid(versionTag(current)) {
		return nil, errors.New("无法比较当前工具与正式发布的版本号。")
	}
	r := &Release{Version: release.Tag, Notes: release.Body, Available: semver.Compare(release.Tag, versionTag(current)) > 0}
	if !r.Available {
		return r, nil
	}
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
		p.report("正在下载更新", fmt.Sprintf("%.1f / %.1f MB（%.0f%%）", float64(p.received)/1e6, float64(p.size)/1e6, 100*float64(p.received)/float64(p.size)))
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
