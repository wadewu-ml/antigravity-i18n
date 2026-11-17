//go:build !windows

package core

import (
	"errors"
	"os/exec"
	"path/filepath"
)

// hideWindow is a no-op outside Windows.
func hideWindow(cmd *exec.Cmd) {}

func isAntigravityRunning(p *paths) (bool, error) {
	if !fileExists(filepath.Join(p.AppDir, antigravityExe)) {
		return false, nil
	}
	return false, errors.New("Antigravity process control requires Windows")
}

func stopAntigravity(p *paths, _ bool, _ int) (bool, error) {
	return isAntigravityRunning(p)
}
