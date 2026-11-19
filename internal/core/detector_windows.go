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

func processExited(handle windows.Handle) (bool, error) {
	state, err := windows.WaitForSingleObject(handle, 0)
	return state == windows.WAIT_OBJECT_0, err
}

// Use a wait-only handle when Windows refuses the image query access. A
// process we cannot identify still blocks the operation while it is alive.
func processIDExited(pid uint32) bool {
	handle, err := windows.OpenProcess(windows.SYNCHRONIZE, false, pid)
	if errors.Is(err, windows.ERROR_INVALID_PARAMETER) {
		return true
	}
	if err != nil {
		return false
	}
	defer windows.CloseHandle(handle)
	exited, err := processExited(handle)
	return err == nil && exited
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
		handle, queryErr := windows.OpenProcess(windows.PROCESS_QUERY_LIMITED_INFORMATION|windows.SYNCHRONIZE, false, entry.ProcessID)
		if errors.Is(queryErr, windows.ERROR_INVALID_PARAMETER) ||
			(errors.Is(queryErr, windows.ERROR_ACCESS_DENIED) && processIDExited(entry.ProcessID)) {
			continue // The process has already exited.
		}
		if queryErr != nil {
			closeProcesses(matched)
			return nil, fmt.Errorf("identifying Antigravity process %d: %w", entry.ProcessID, queryErr)
		}
		exited, waitErr := processExited(handle)
		if waitErr != nil || exited {
			_ = windows.CloseHandle(handle)
			if waitErr != nil {
				closeProcesses(matched)
				return nil, fmt.Errorf("checking Antigravity process %d: %w", entry.ProcessID, waitErr)
			}
			continue
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
		if queryErr != nil && !os.IsNotExist(queryErr) && !errors.Is(queryErr, windows.ERROR_INVALID_PARAMETER) {
			// The process can exit between the signal check and the image query.
			exited, err := processExited(handle)
			if err != nil || !exited {
				_ = windows.CloseHandle(handle)
				closeProcesses(matched)
				return nil, fmt.Errorf("identifying Antigravity process %d: %w", entry.ProcessID, queryErr)
			}
		}
		_ = windows.CloseHandle(handle)
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
	defer func() { closeProcesses(processes) }()

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
	runErr := cmd.Run() // A selected process may have exited before taskkill reached it.

	for {
		// The handles identify the processes we already checked. Querying their
		// image paths again during termination can fail before Windows drops the PID.
		allExited := true
		for _, process := range processes {
			exited, err := processExited(process.handle)
			if err != nil {
				return true, fmt.Errorf("waiting for Antigravity process %d to exit: %w", process.pid, err)
			}
			if !exited {
				allExited = false
				break
			}
		}
		if allExited {
			closeProcesses(processes)
			processes = nil
			running, err := isAntigravityRunning(p)
			if err == nil && running {
				err = errors.New("Antigravity restarted during shutdown. Close it and re-run")
			}
			return true, err
		}
		remaining := time.Until(deadline)
		if remaining <= 0 {
			message := "Antigravity did not exit before the timeout. Close it manually or explicitly choose forced shutdown and retry"
			if force {
				message = "Antigravity is still running after forced shutdown. Close it and retry"
			}
			if runErr != nil {
				return true, fmt.Errorf("%s (taskkill: %w)", message, runErr)
			}
			return true, errors.New(message)
		}
		time.Sleep(min(remaining, gracefulPollIntervalMs*time.Millisecond))
	}
}
