package core

import (
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
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	path := filepath.ToSlash(page)
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	pageURL := (&url.URL{Scheme: "file", Path: path}).String()
	cmd := exec.CommandContext(ctx, browser, "--headless", "--disable-gpu", "--no-first-run",
		"--no-default-browser-check", "--user-data-dir="+filepath.Join(dir, "profile"),
		"--virtual-time-budget=1000", "--dump-dom", pageURL)
	hideWindow(cmd)
	output, err := cmd.CombinedOutput()
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
