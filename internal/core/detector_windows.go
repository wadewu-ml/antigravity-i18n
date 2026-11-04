//go:build windows

package core

import (
	"os/exec"
	"syscall"
)

// hideWindow keeps helper commands from flashing a console window.
func hideWindow(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true, CreationFlags: 0x08000000}
}
