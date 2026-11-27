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
	"sync"
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

type browserOutput struct {
	mu  sync.Mutex
	buf bytes.Buffer
}

func (b *browserOutput) Write(data []byte) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.buf.Write(data)
}

func (b *browserOutput) snapshot() []byte {
	b.mu.Lock()
	defer b.mu.Unlock()
	return bytes.Clone(b.buf.Bytes())
}

var browserResult = regexp.MustCompile(`data-ag-test-result="([A-Za-z0-9+/=]+)"`)

// dumpBrowserDOM returns Chrome's dumped DOM. Some headless builds print the
// document and then never exit, so the dump is collected until the test result
// is present or the process ends.
func dumpBrowserDOM(ctx context.Context, cmd *exec.Cmd) ([]byte, error) {
	var buf browserOutput
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	done := make(chan error, 1)
	go func() { done <- cmd.Wait() }()
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	for {
		if browserResult.Match(buf.snapshot()) {
			if cmd.Process != nil {
				_ = cmd.Process.Kill()
			}
			<-done
			return buf.snapshot(), nil
		}
		select {
		case err := <-done:
			if err != nil {
				return buf.snapshot(), err
			}
			return buf.snapshot(), nil
		case <-ctx.Done():
			if cmd.Process != nil {
				_ = cmd.Process.Kill()
			}
			<-done
			return buf.snapshot(), ctx.Err()
		case <-ticker.C:
		}
	}
}

func TestTranslationInBrowser(t *testing.T) {
	store := newTestService(t).store
	for _, example := range []struct{ code, fixture string }{
		{DefaultLocale, "testdata/engine.html"},
		{"de", "testdata/engine-repeat.html"},
		{"tr", "testdata/engine-tr-vi.html"},
		{"vi", "testdata/engine-tr-vi.html"},
		{"ar", "testdata/engine-ar.html"},
		{"it", "testdata/engine-it.html"},
		{"id", "testdata/engine-id.html"},
		{"pl", "testdata/engine-pl.html"},
		{"cs", "testdata/engine-cs-hu-bg.html"},
		{"hu", "testdata/engine-cs-hu-bg.html"},
		{"bg", "testdata/engine-cs-hu-bg.html"},
	} {
		t.Run(example.code, func(t *testing.T) {
			loc, err := store.LoadLocale(example.code)
			if err != nil {
				t.Fatal(err)
			}
			if example.fixture == "testdata/engine-repeat.html" {
				// Keep a translated value that is also a lookup key so the fixture
				// still exercises the mutation guard when real translations change.
				loc.Text["Test label"] = "Shared label"
				loc.Text["Shared label"] = "Final label"
			}
			engine, err := store.buildPreloadFragment(loc)
			if err != nil {
				t.Fatal(err)
			}
			html := readTestFile(t, example.fixture)
			runBrowserChecks(t, strings.Replace(html, "<!-- ENGINE -->", "<script>"+engine+"</script>", 1))
		})
	}
}

func TestRetryMessageInEveryLanguage(t *testing.T) {
	store := newTestService(t).store
	for _, code := range store.ListLocales() {
		t.Run(code, func(t *testing.T) {
			loc, err := store.LoadLocale(code)
			if err != nil {
				t.Fatal(err)
			}
			script, err := store.buildPreloadFragment(loc)
			if err != nil {
				t.Fatal(err)
			}
			page := readTestFile(t, "testdata/retry-message.html")
			runBrowserChecks(t, strings.Replace(page, "<!-- ENGINE -->", "<script>"+script+"</script>", 1))
		})
	}
}

func TestEffortSliderInBrowser(t *testing.T) {
	store := newTestService(t).store
	for _, code := range []string{DefaultLocale, "it", "ar", "id", "pl", "cs", "hu", "bg"} {
		t.Run(code, func(t *testing.T) {
			loc, err := store.LoadLocale(code)
			if err != nil {
				t.Fatal(err)
			}
			patch, err := store.buildLocalizationFragments(loc)
			if err != nil {
				t.Fatal(err)
			}
			if err := store.addEffortSlider(&patch, loc); err != nil {
				t.Fatal(err)
			}
			html := readTestFile(t, "testdata/effort-slider.html")
			runBrowserChecks(t, strings.Replace(html, "<!-- ENGINE -->", "<script>"+patch.preload+"</script>", 1))
		})
	}
}

func TestSetupLanguageInBrowser(t *testing.T) {
	runSetupBrowserChecks(t, "testdata/setup-language.js")
}

func TestSetupLayoutInBrowser(t *testing.T) {
	runSetupBrowserChecks(t, "testdata/setup-layout.js")
}

func TestApplyShutdownInBrowser(t *testing.T) {
	runSetupBrowserChecks(t, "testdata/shutdown.js")
}

func TestStatusSelectionInBrowser(t *testing.T) {
	runSetupBrowserChecks(t, "testdata/status-selection.js")
}

func TestAutoUpdateInBrowser(t *testing.T) {
	for _, scenario := range []string{"available", "installing", "cached", "disabled", "offline", "limited", "upgraded"} {
		t.Run(scenario, func(t *testing.T) {
			runSetupBrowserChecks(t, "testdata/auto-update.js", "const scenario = '"+scenario+"';")
		})
	}
}

func runSetupBrowserChecks(t *testing.T, script string, setup ...string) {
	t.Helper()
	html := readTestFile(t, "../../frontend/dist/index.html")
	style := readTestFile(t, "../../frontend/dist/style.css")
	html = strings.Replace(html, `<link rel="stylesheet" href="style.css">`, "<style>"+style+"</style>", 1)
	localeData := readTestFile(t, "../../frontend/dist/setup-locales.js")
	html = strings.Replace(html, `<script src="setup-locales.js"></script>`, "<script>"+localeData+"</script>", 1)
	app := readTestFile(t, "../../frontend/dist/app.js")
	checks := strings.Join(setup, "\n") + "\n" + readTestFile(t, script)
	scripts := "<script>" + checks + "</script><script>" + app + "</script>"
	runBrowserChecks(t, strings.Replace(html, `<script src="app.js"></script>`, scripts, 1), "--force-device-scale-factor=1", "--window-size=760,700")
}

func readTestFile(t *testing.T, path string) string {
	t.Helper()
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return string(content)
}

const browserAssertions = `<script>
const failures = [];
let checks = 0;
function check(label, actual, expected) {
  checks++;
  if (actual !== expected) failures.push(label + ': expected ' + expected + ', got ' + actual);
}
function settle(wait = 0) { return new Promise(resolve => setTimeout(resolve, wait)); }
function verify(run) {
  addEventListener('load', async () => {
    try { await run(); } catch (error) { failures.push(String(error)); }
    const result = btoa(unescape(encodeURIComponent(JSON.stringify({ checks, failures }))));
    document.documentElement.setAttribute('data-ag-test-result', result);
  }, { once: true });
}
</script>`

func runBrowserChecks(t *testing.T, html string, browserArgs ...string) {
	t.Helper()
	browser := engineTestBrowser(t)
	dir := t.TempDir()
	page := filepath.Join(dir, "engine.html")
	source := strings.Replace(html, "</head>", browserAssertions+"</head>", 1)
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
	args := []string{"--headless=new", "--disable-gpu", "--no-sandbox",
		"--disable-dev-shm-usage", "--no-first-run", "--no-default-browser-check",
		"--user-data-dir=" + filepath.Join(dir, "profile"),
		"--virtual-time-budget=8000", "--dump-dom"}
	args = append(args, browserArgs...)
	cmd := exec.CommandContext(ctx, browser, append(args, pageURL)...)
	hideWindow(cmd)
	output, err := dumpBrowserDOM(ctx, cmd)
	if err != nil {
		t.Fatalf("headless browser: %v\n%s", err, output)
	}
	match := browserResult.FindSubmatch(output)
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
	if result.Checks == 0 || len(result.Failures) > 0 {
		t.Fatalf("%d DOM checks, failures: %v", result.Checks, result.Failures)
	}
	t.Logf("%d DOM checks passed", result.Checks)
}
