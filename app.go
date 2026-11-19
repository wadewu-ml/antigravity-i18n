package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	"antigravity-i18n/internal/core"

	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

const confirmationRequired = "请先确认是否继续。"

// App binds the Go core to the Wails frontend. The GUI only reaches the core
// through this type, so a window click and a CLI command perform the same work.
type App struct {
	ctx      context.Context
	svc      *core.Service
	appDir   string
	mu       sync.Mutex
	opMu     sync.Mutex
	logs     []string
	progress func(stage, detail string)
}

// NewApp builds the GUI application around the shared locale store.
func NewApp(svc *core.Service, initialDir string) *App {
	return &App{svc: svc, appDir: initialDir}
}

// startup captures the context the runtime needs to emit events.
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.progress = func(stage, detail string) {
		wailsRuntime.EventsEmit(ctx, "progress", map[string]string{
			"stage":  stage,
			"detail": detail,
		})
	}
	// The sink is installed once here rather than at each call site, so an
	// action taken before startup has already been set cannot go silent.
	a.svc.Progress = a.progress
	a.svc.Progress("就绪", "正在检测 Antigravity")
}

func (a *App) appendLog(msg string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.logs = append(a.logs, msg)
	if len(a.logs) > 400 {
		a.logs = a.logs[len(a.logs)-400:]
	}
}

func (a *App) currentDir() string {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.appDir
}

// GetVersion reports the tool version shown in the window chrome.
func (a *App) GetVersion() string { return core.Version }

// GetStatus reports the localization state of the selected installation.
func (a *App) GetStatus() (*core.Status, error) {
	st, err := a.svc.GetStatus(a.currentDir())
	if err != nil {
		return nil, err
	}
	if st != nil && st.AppDir != "" {
		a.mu.Lock()
		a.appDir = st.AppDir
		a.mu.Unlock()
	}
	return st, nil
}

// GetLocales lists every bundled language pack.
func (a *App) GetLocales() ([]core.LocaleInfo, error) {
	return a.svc.DescribeLocales()
}

// SelectAppDir opens a native folder picker.
func (a *App) SelectAppDir() (string, error) {
	if a.ctx == nil {
		return "", errors.New("app context not initialized")
	}
	selected, err := wailsRuntime.OpenDirectoryDialog(a.ctx, wailsRuntime.OpenDialogOptions{
		Title: "Select the Antigravity installation folder (the one holding Antigravity.exe)",
	})
	if err != nil {
		return "", err
	}
	return selected, nil
}

// SetAppDir targets another installation and returns the refreshed status.
func (a *App) SetAppDir(path string) (*core.Status, error) {
	if path == "" {
		return nil, errors.New("no directory selected")
	}
	st, err := a.svc.GetStatus(filepath.Clean(path))
	if err != nil {
		return nil, err
	}
	a.mu.Lock()
	a.appDir = st.AppDir
	a.mu.Unlock()
	return st, nil
}

// ApplyLocale installs a language pack, replacing whatever patch is present.
func (a *App) ApplyLocale(locale string, confirmed, force bool) (*core.ApplyResult, error) {
	return a.applyCustomization(locale, nil, confirmed, force)
}

// ApplyCustomization installs a language pack with the selected features.
func (a *App) ApplyCustomization(locale string, sessionMetrics, confirmed, force bool) (*core.ApplyResult, error) {
	return a.applyCustomization(locale, &sessionMetrics, confirmed, force)
}

func (a *App) applyCustomization(locale string, sessionMetrics *bool, confirmed, force bool) (*core.ApplyResult, error) {
	if !confirmed {
		return nil, errors.New(confirmationRequired)
	}
	if !a.opMu.TryLock() {
		return nil, errors.New("正在执行另一项语言操作。")
	}
	defer a.opMu.Unlock()
	result, err := a.svc.ApplyLocale(core.Options{
		AppDir:         a.currentDir(),
		Locale:         locale,
		Restart:        true,
		Force:          force,
		SessionMetrics: sessionMetrics,
	})
	for _, note := range a.svc.Notes() {
		a.appendLog("note: " + note)
	}
	if err != nil {
		a.appendLog("error: " + err.Error())
		return nil, err
	}
	a.appendLog(fmt.Sprintf("applied %s (%s)", result.LocaleName, result.Locale))
	return result, nil
}

// RestoreLocale puts the official English build back. Restore stops the process
// exactly like apply does, so it demands the same confirmation.
func (a *App) RestoreLocale(confirmed, force bool) (*core.RestoreResult, error) {
	if !confirmed {
		return nil, errors.New(confirmationRequired)
	}
	if !a.opMu.TryLock() {
		return nil, errors.New("正在执行另一项语言操作。")
	}
	defer a.opMu.Unlock()
	result, err := a.svc.RestoreOfficial(core.Options{
		AppDir:  a.currentDir(),
		Restart: true,
		Force:   force,
	})
	for _, note := range a.svc.Notes() {
		a.appendLog("note: " + note)
	}
	if err != nil {
		a.appendLog("error: " + err.Error())
		return nil, err
	}
	a.appendLog("restored the official English build")
	return result, nil
}

// LaunchAntigravity starts the application after a successful patch.
func (a *App) LaunchAntigravity() error {
	dir := a.currentDir()
	if dir == "" {
		return errors.New("no Antigravity installation selected")
	}
	exe := filepath.Join(dir, "Antigravity.exe")
	if _, err := os.Stat(exe); err != nil {
		return fmt.Errorf("Antigravity.exe was not found in %s", dir)
	}
	cmd := exec.Command(exe)
	if err := cmd.Start(); err != nil {
		return err
	}
	_ = cmd.Process.Release()
	return nil
}

// GetLogs returns the diagnostics collected in this session.
func (a *App) GetLogs() []string {
	a.mu.Lock()
	defer a.mu.Unlock()
	out := make([]string, len(a.logs))
	copy(out, a.logs)
	return out
}
