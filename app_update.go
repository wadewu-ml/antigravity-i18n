package main

import (
	"context"
	"errors"
	"time"

	"antigravity-i18n/internal/core"
	"antigravity-i18n/internal/selfupdate"
	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) CheckForUpdate() (*selfupdate.Release, error) {
	if !a.opMu.TryLock() {
		return nil, errors.New("正在执行另一项操作。")
	}
	defer a.opMu.Unlock()
	release, err := selfupdate.Check(a.ctx, core.Version)
	a.offered = release
	return release, err
}

func (a *App) OpenReleases() {
	wailsRuntime.BrowserOpenURL(a.ctx, selfupdate.Repository+"/releases")
}

func (a *App) InstallUpdate(version string) error {
	if !a.opMu.TryLock() {
		return errors.New("正在执行另一项操作。")
	}
	if a.offered == nil || !a.offered.Available || a.offered.Version != version {
		a.opMu.Unlock()
		return errors.New("请先检查并确认更新版本。")
	}
	ctx, cancel := context.WithTimeout(a.ctx, 10*time.Minute)
	defer cancel()
	a.mu.Lock()
	a.installing, a.cancelDownload = true, cancel
	a.mu.Unlock()

	err := a.offered.Install(ctx, core.Version, a.currentDir(), a.progress, func() error {
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
