package core

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"
)

// Use an isolated headless browser so TreeWalker, computed styles and mutation
// delivery are exercised by a real DOM. No npm packages or running app required.
// ANTIGRAVITY_TEST_BROWSER can name a Chromium executable on other platforms.
func engineTestBrowser(t *testing.T) string {
	t.Helper()
	if browser := os.Getenv("ANTIGRAVITY_TEST_BROWSER"); browser != "" {
		return browser
	}
	for _, name := range []string{"chromium", "chromium-browser", "google-chrome", "msedge"} {
		if browser, err := exec.LookPath(name); err == nil {
			return browser
		}
	}
	for _, root := range []string{os.Getenv("PROGRAMFILES"), os.Getenv("PROGRAMFILES(X86)"), os.Getenv("LOCALAPPDATA")} {
		if root == "" {
			continue
		}
		for _, rel := range []string{"Google/Chrome/Application/chrome.exe", "Microsoft/Edge/Application/msedge.exe"} {
			browser := filepath.Join(root, filepath.FromSlash(rel))
			if info, err := os.Stat(browser); err == nil && !info.IsDir() {
				return browser
			}
		}
	}
	t.Skip("set ANTIGRAVITY_TEST_BROWSER to a Chromium executable to run DOM regression tests")
	return ""
}

// dumpBrowserDOM returns Chrome's dumped DOM. Some headless builds print the
// document and then never exit, so the dump is collected until the test result
// is present or the process ends.
func dumpBrowserDOM(ctx context.Context, cmd *exec.Cmd) ([]byte, error) {
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	done := make(chan error, 1)
	go func() { done <- cmd.Wait() }()
	result := regexp.MustCompile(`data-ag-test-result="[A-Za-z0-9+/=]+"`)
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	for {
		if result.Match(buf.Bytes()) {
			if cmd.Process != nil {
				_ = cmd.Process.Kill()
			}
			<-done
			return buf.Bytes(), nil
		}
		select {
		case err := <-done:
			if err != nil {
				return buf.Bytes(), err
			}
			return buf.Bytes(), nil
		case <-ctx.Done():
			if cmd.Process != nil {
				_ = cmd.Process.Kill()
			}
			<-done
			return buf.Bytes(), ctx.Err()
		case <-ticker.C:
		}
	}
}

func TestEnginePreservesContentAndSentenceFragments(t *testing.T) {
	browser := engineTestBrowser(t)
	assets := os.DirFS("../..")
	store, err := NewLocaleStore(assets, assets, "assets/locales", "assets/patches")
	if err != nil {
		t.Fatal(err)
	}
	loc, err := store.LoadLocale(DefaultLocale)
	if err != nil {
		t.Fatal(err)
	}
	engine, err := store.buildPreloadFragment(loc)
	if err != nil {
		t.Fatal(err)
	}
	fixture, err := os.ReadFile("testdata/engine.html")
	if err != nil {
		t.Fatal(err)
	}
	dir := t.TempDir()
	page := filepath.Join(dir, "engine.html")
	source := strings.Replace(string(fixture), "<!-- ENGINE -->", "<script>"+engine+"</script>", 1)
	if err := os.WriteFile(page, []byte(source), 0o600); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()
	path := filepath.ToSlash(page)
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	pageURL := (&url.URL{Scheme: "file", Path: path}).String()
	// Chrome in this kind of container prints the DOM and then stays up on a
	// missing system bus, so the result is read from stdout instead of waiting
	// for the process to exit. --no-sandbox is required for that container.
	cmd := exec.CommandContext(ctx, browser, "--headless=new", "--disable-gpu", "--no-sandbox",
		"--disable-dev-shm-usage", "--no-first-run", "--no-default-browser-check",
		"--user-data-dir="+filepath.Join(dir, "profile"),
		"--virtual-time-budget=3000", "--dump-dom", pageURL)
	hideWindow(cmd)
	output, err := dumpBrowserDOM(ctx, cmd)
	if err != nil {
		t.Fatalf("headless browser: %v\n%s", err, output)
	}
	match := regexp.MustCompile(`data-ag-test-result="([A-Za-z0-9+/=]+)"`).FindSubmatch(output)
	if match == nil {
		t.Fatalf("browser did not finish the engine checks:\n%s", output)
	}
	raw, err := base64.StdEncoding.DecodeString(string(match[1]))
	if err != nil {
		t.Fatal(err)
	}
	var result struct {
		Checks   int      `json:"checks"`
		Failures []string `json:"failures"`
	}
	if err := json.Unmarshal(raw, &result); err != nil {
		t.Fatal(err)
	}
	if result.Checks < 30 || len(result.Failures) > 0 {
		t.Fatalf("%d DOM checks, failures: %v", result.Checks, result.Failures)
	}
	t.Logf("%d DOM checks passed", result.Checks)
}

func TestChinesePacksPreserveTurbo(t *testing.T) {
	assets := os.DirFS("../..")
	store, err := NewLocaleStore(assets, assets, "assets/locales", "assets/patches")
	if err != nil {
		t.Fatal(err)
	}
	for _, code := range []string{"zh-CN", "zh-Hant"} {
		loc, err := store.LoadLocale(code)
		if err != nil {
			t.Fatal(err)
		}
		if loc.Text["Turbo"] != "Turbo" || loc.Text["Turbo mode"] != "Turbo 模式" {
			t.Errorf("%s should preserve the permission preset name Turbo", code)
		}
	}
}
