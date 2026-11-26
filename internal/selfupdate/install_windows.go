package selfupdate

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"golang.org/x/sys/windows"
)

// The named handle exists only while the update or cleanup is in progress.
// Windows releases it even if the process is interrupted.
func lockExecutable(exe string) (func(), error) {
	key := sha256.Sum256([]byte(strings.ToLower(exe)))
	name, err := windows.UTF16PtrFromString(fmt.Sprintf("Global\\regravity-update-%x", key))
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

// Keep the old GUI and its rollback path alive until the new frontend calls
// UpdateWindowReady. The process handle also wakes the wait on an early exit.
func startAndWaitReady(cmd *exec.Cmd, id string, timeout time.Duration) error {
	name := "Local\\Regravity.Ready." + id
	eventName, err := windows.UTF16PtrFromString(name)
	if err != nil {
		return err
	}
	event, err := windows.CreateEvent(nil, 1, 0, eventName)
	if event != 0 {
		defer windows.CloseHandle(event)
	}
	if err != nil {
		return err
	}
	cmd.Env = append(cmd.Environ(), readyEnv+"="+name)
	if err := cmd.Start(); err != nil {
		return err
	}
	var result uint32
	var waitErr error
	err = cmd.Process.WithHandle(func(handle uintptr) {
		result, waitErr = windows.WaitForMultipleObjects([]windows.Handle{windows.Handle(handle), event}, false, uint32(timeout.Milliseconds()))
	})
	if err == nil && waitErr == nil && result == windows.WAIT_OBJECT_0+1 {
		return cmd.Process.Release()
	}
	// Wait for Windows to release the mapped executable before restoring it.
	_ = cmd.Process.Kill()
	_ = cmd.Wait()
	if err := errors.Join(err, waitErr); err != nil {
		return err
	}
	if result == uint32(windows.WAIT_TIMEOUT) {
		return errors.New("等待新版窗口就绪超时。")
	}
	return errors.New("新版程序在窗口就绪前退出。")
}

func signalReady() error {
	name := os.Getenv(readyEnv)
	if name == "" {
		return nil
	}
	_ = os.Unsetenv(readyEnv)
	eventName, err := windows.UTF16PtrFromString(name)
	if err != nil {
		return err
	}
	event, err := windows.OpenEvent(windows.EVENT_MODIFY_STATE, false, eventName)
	if err != nil {
		return err
	}
	defer windows.CloseHandle(event)
	return windows.SetEvent(event)
}
