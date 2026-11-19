package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"

	"antigravity-i18n/internal/core"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

// runWails opens the application window. The frontend is plain static files
// served from an embedded filesystem, so no Node toolchain is needed at build
// time and the binary stays a single self-contained executable.
func runWails(svc *core.Service) {
	distFS, err := fs.Sub(frontendFS, "frontend/dist")
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	app := NewApp(svc, "")
	if err := wails.Run(&options.App{
		Title:     "Antigravity 多语言工具",
		Width:     860,
		Height:    760,
		MinWidth:  640,
		MinHeight: 620,
		AssetServer: &assetserver.Options{
			Assets: distFS,
		},
		BackgroundColour: &options.RGBA{R: 250, G: 251, B: 253, A: 255},
		OnStartup:        app.startup,
		ErrorFormatter:   formatAppError,
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    false,
		},
	}); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func formatAppError(err error) any {
	var rollbackErr *core.RollbackError
	if errors.Is(err, core.ErrGracefulShutdownTimeout) && !errors.As(err, &rollbackErr) {
		// Wails passes the formatted value through the JavaScript Error constructor.
		return "shutdown_timeout: " + err.Error()
	}
	return err.Error()
}
