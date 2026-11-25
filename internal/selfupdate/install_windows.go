package selfupdate

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"syscall"

	"golang.org/x/sys/windows"
)

// The named handle exists only while the update or cleanup is in progress.
// Windows releases it even if the process is interrupted.
func lockExecutable(exe string) (func(), error) {
	key := sha256.Sum256([]byte(strings.ToLower(exe)))
	name, err := windows.UTF16PtrFromString(fmt.Sprintf("Global\\antigravity-i18n-update-%x", key))
	if err != nil {
		return nil, err
	}
	h, err := windows.CreateMutex(nil, false, name)
	if errors.Is(err, windows.ERROR_ALREADY_EXISTS) {
		if h != 0 {
			_ = windows.CloseHandle(h)
		}
		return nil, errors.New("另一个窗口正在更新这个工具，请稍后重试。")
	}
	if err != nil {
		return nil, err
	}
	return func() { _ = windows.CloseHandle(h) }, nil
}

func hideProbeWindow(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
}

// Omitting REPLACE_EXISTING preserves any existing destination on rollback.
func moveNewName(from, to string) error {
	src, err := windows.UTF16PtrFromString(from)
	if err != nil {
		return err
	}
	dst, err := windows.UTF16PtrFromString(to)
	if err != nil {
		return err
	}
	return windows.MoveFileEx(src, dst, windows.MOVEFILE_WRITE_THROUGH)
}
