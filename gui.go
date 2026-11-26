//go:build !cli

package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"regravity/internal/core"
	"regravity/internal/selfupdate"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

const guiHelp = "  gui        Open the graphical interface (default when no command is given)\n"

const guiWindowClass = "WadeWu.Regravity.MainWindow"

// runWails opens the application window. The frontend is plain static files
// served from an embedded filesystem, so no Node toolchain is needed at build
// time and the binary stays a single self-contained executable.
func runWails(svc *core.Service) {
	initialDir, restarting := os.LookupEnv(selfupdate.RestartDirEnv)
	_ = os.Unsetenv(selfupdate.RestartDirEnv)
	release, err := enterGUI(restarting)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	if release == nil {
		return // The existing window handled this launch.
	}
	defer release()

	distFS, err := fs.Sub(frontendFS, "frontend/dist")
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	profile, err := os.UserConfigDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error locating the settings directory:", err)
		os.Exit(1)
	}
	app := NewApp(svc, initialDir)
	if err := wails.Run(&options.App{
		Title:     "Regravity",
		Width:     760,
		Height:    580,
		MinWidth:  640,
		MinHeight: 520,
		AssetServer: &assetserver.Options{
			Assets: distFS,
		},
		BackgroundColour: &options.RGBA{R: 244, G: 246, B: 249, A: 255},
		OnStartup:        app.startup,
		OnBeforeClose:    app.beforeClose,
		ErrorFormatter:   formatAppError,
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			// The previous Windows release stored update preferences in this profile.
			WebviewUserDataPath:  filepath.Join(profile, "antigravity-i18n.exe"),
			WindowClassName:      guiWindowClass,
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    false,
		},
	}); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
