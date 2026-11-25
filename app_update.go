package main

import (
	"context"
	"errors"
	"time"

	"antigravity-i18n/internal/core"
	"antigravity-i18n/internal/selfupdate"
	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

type UpdateCheckResult struct {
	Release *selfupdate.Release `json:"release,omitempty"`
	Error   string              `json:"error,omitempty"`
	RetryAt int64               `json:"retryAt,omitempty"`
}

func (a *App) CheckForUpdate() UpdateCheckResult {
	// Checking is read-only. Apply and Restore must stay available during network I/O.
	release, err := selfupdate.Check(a.ctx, core.Version)
	if err != nil {
		result := UpdateCheckResult{Error: err.Error()}
		var limited *selfupdate.RateLimitError
		if errors.As(err, &limited) {
			result.RetryAt = limited.RetryAt.UnixMilli()
		}
		return result
	}
	a.mu.Lock()
	a.offered = release
	a.mu.Unlock()
	return UpdateCheckResult{Release: release}
}

func (a *App) OpenReleases() {
	wailsRuntime.BrowserOpenURL(a.ctx, selfupdate.Repository+"/releases")
}

func (a *App) InstallUpdate(version string) error {
	if !a.opMu.TryLock() {
		return errors.New("正在执行另一项操作。")
	}
	a.mu.Lock()
	offered := a.offered
	a.mu.Unlock()
	if offered == nil || !offered.Available || offered.Version != version {
		a.opMu.Unlock()
		return errors.New("请先检查并确认更新版本。")
	}
	ctx, cancel := context.WithTimeout(a.ctx, 10*time.Minute)
	defer cancel()
	a.mu.Lock()
	a.installing, a.cancelDownload = true, cancel
	a.mu.Unlock()

	err := offered.Install(ctx, core.Version, a.currentDir(), a.progress, func() error {
		a.mu.Lock()
		defer a.mu.Unlock()
		if err := ctx.Err(); err != nil {
			return err
		}
		a.cancelDownload = nil
		return nil
	})

	a.mu.Lock()
	closeWindow := err == nil || a.closeAfterUpdate
	a.installing, a.closeAfterUpdate, a.cancelDownload = false, false, nil
	a.mu.Unlock()
	if closeWindow {
		// Keep other operations out while Wails tears down the old window.
		wailsRuntime.Quit(a.ctx)
	} else {
		a.opMu.Unlock()
	}
	if errors.Is(err, context.Canceled) {
		return errors.New("已取消更新。")
	}
	return err
}

func (a *App) CancelUpdate() {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.cancelDownload != nil {
		a.cancelDownload()
	}
}

func (a *App) beforeClose(context.Context) bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.installing {
		if a.cancelDownload != nil {
			a.closeAfterUpdate = true
			a.cancelDownload()
		}
		return true
	}
	return false
}

func (a *App) UpdateWindowReady() {
	a.readyOnce.Do(func() {
		go func() {
			if err := selfupdate.WindowReady(core.Version); err != nil {
				a.appendLog("更新清理：" + err.Error())
			}
		}()
	})
}
