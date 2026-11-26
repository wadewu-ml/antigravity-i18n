//go:build !cli

package main

import (
	"errors"
	"fmt"
	"os"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	user32              = windows.NewLazySystemDLL("user32.dll")
	findWindow          = user32.NewProc("FindWindowW")
	isIconic            = user32.NewProc("IsIconic")
	showWindowAsync     = user32.NewProc("ShowWindowAsync")
	setForegroundWindow = user32.NewProc("SetForegroundWindow")
)

// A named handle elects one GUI per interactive session, including when two
// launches race before Wails has created a window. Windows closes it on exit.
func enterGUI(restarting bool) (release func(), err error) {
	mutexName, _ := windows.UTF16PtrFromString("Local\\" + guiWindowClass)
	className, _ := windows.UTF16PtrFromString(guiWindowClass)
	var updaterPID uint32
	if restarting {
		updaterPID = uint32(os.Getppid())
	}
	deadline := time.Now().Add(15 * time.Second)
	for {
		h, err := windows.CreateMutex(nil, false, mutexName)
		if err == nil {
			return func() { windows.CloseHandle(h) }, nil
		}
		if !errors.Is(err, windows.ERROR_ALREADY_EXISTS) {
			return nil, fmt.Errorf("无法打开窗口锁：%w", err)
		}

		hwnd, _, _ := findWindow.Call(uintptr(unsafe.Pointer(className)), 0)
		if windows.IsWindowVisible(windows.HWND(hwnd)) {
			var pid uint32
			if _, err := windows.GetWindowThreadProcessId(windows.HWND(hwnd), &pid); err == nil {
				if restarting && pid == updaterPID {
					// Both processes retain the instance handle during the handoff.
					return func() { windows.CloseHandle(h) }, nil
				}
				windows.CloseHandle(h)
				// This newly launched process can activate the window on behalf of
				// the user who just double-clicked it. Preserve maximised windows.
				if minimised, _, _ := isIconic.Call(hwnd); minimised != 0 {
					const swRestore = 9
					showWindowAsync.Call(hwnd, swRestore)
				}
				setForegroundWindow.Call(hwnd)
				return nil, nil
			}
		}
		windows.CloseHandle(h)
		if time.Now().After(deadline) {
			return nil, errors.New("已打开的工具窗口暂时没有响应。")
		}
		time.Sleep(50 * time.Millisecond)
	}
}
