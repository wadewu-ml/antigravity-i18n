//go:build windows

package core

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

// hideWindow keeps helper commands from flashing a console window.
func hideWindow(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true, CreationFlags: 0x08000000}
}

type antigravityProcess struct {
	pid    uint32
	handle windows.Handle
}

func closeProcesses(processes []antigravityProcess) {
	for _, process := range processes {
		_ = windows.CloseHandle(process.handle)
	}
}

// matchingProcesses retains handles until the caller finishes using the PIDs.
func matchingProcesses(p *paths) ([]antigravityProcess, error) {
	target, err := filepath.Abs(filepath.Join(p.AppDir, antigravityExe))
	if err != nil {
		return nil, err
	}
	target, err = filepath.EvalSymlinks(target)
	if os.IsNotExist(err) {
		// A standalone archive has no application process to stop.
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	snapshot, err := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPPROCESS, 0)
	if err != nil {
		return nil, fmt.Errorf("listing Antigravity processes: %w", err)
	}
	defer windows.CloseHandle(snapshot)

	var matched []antigravityProcess
	var entry windows.ProcessEntry32
	entry.Size = uint32(unsafe.Sizeof(entry))
	path := make([]uint16, windows.MAX_LONG_PATH)
	for err = windows.Process32First(snapshot, &entry); err == nil; err = windows.Process32Next(snapshot, &entry) {
		if !strings.EqualFold(windows.UTF16ToString(entry.ExeFile[:]), antigravityExe) {
			continue
		}
		handle, queryErr := windows.OpenProcess(windows.PROCESS_QUERY_LIMITED_INFORMATION, false, entry.ProcessID)
		if errors.Is(queryErr, windows.ERROR_INVALID_PARAMETER) {
			continue // The process has already exited.
		}
		if queryErr != nil {
			closeProcesses(matched)
			return nil, fmt.Errorf("identifying Antigravity process %d: %w", entry.ProcessID, queryErr)
		}
		length := uint32(len(path))
		queryErr = windows.QueryFullProcessImageName(handle, 0, &path[0], &length)
		if queryErr == nil {
			var image string
			image, queryErr = filepath.EvalSymlinks(windows.UTF16ToString(path[:length]))
			if queryErr == nil && image == target {
				matched = append(matched, antigravityProcess{pid: entry.ProcessID, handle: handle})
				continue
			}
		}
		_ = windows.CloseHandle(handle)
		if queryErr != nil && !os.IsNotExist(queryErr) && !errors.Is(queryErr, windows.ERROR_INVALID_PARAMETER) {
			closeProcesses(matched)
			return nil, fmt.Errorf("identifying Antigravity process %d: %w", entry.ProcessID, queryErr)
		}
	}
	if !errors.Is(err, windows.ERROR_NO_MORE_FILES) {
		closeProcesses(matched)
		return nil, fmt.Errorf("listing Antigravity processes: %w", err)
	}
	return matched, nil
}

func isAntigravityRunning(p *paths) (bool, error) {
	processes, err := matchingProcesses(p)
	closeProcesses(processes)
	return len(processes) > 0, err
}

func stopAntigravity(p *paths, force bool, timeoutMs int) (bool, error) {
	processes, err := matchingProcesses(p)
	if err != nil || len(processes) == 0 {
		return false, err
	}
	defer closeProcesses(processes)

	// Include child processes such as the Antigravity language server.
	args := []string{"/T"}
	if force {
		args = append(args, "/F")
	}
	for _, process := range processes {
		args = append(args, "/PID", strconv.FormatUint(uint64(process.pid), 10))
	}
	deadline := time.Now().Add(time.Duration(timeoutMs) * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	cmd := exec.CommandContext(ctx, "taskkill.exe", args...)
	hideWindow(cmd)
	_ = cmd.Run() // Poll the installation in case a selected process already exited.

	for {
		running, err := isAntigravityRunning(p)
		if err != nil || !running {
			return true, err
		}
		remaining := time.Until(deadline)
		if remaining <= 0 {
			if force {
				return true, errors.New("Antigravity is still running after forced shutdown. Close it and retry")
			}
			return true, errors.New("Antigravity did not exit before the timeout. Close it manually or explicitly choose forced shutdown and retry")
		}
		time.Sleep(min(remaining, gracefulPollIntervalMs*time.Millisecond))
	}
}
