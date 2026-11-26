//go:build !windows

package selfupdate

import (
	"errors"
	"os/exec"
	"time"
)

var errWindowsOnly = errors.New("工具自更新仅支持 Windows。")

func lockExecutable(string) (func(), error) { return nil, errWindowsOnly }
func hideProbeWindow(*exec.Cmd)             {}
func moveNewName(string, string) error      { return errWindowsOnly }

func startAndWaitReady(*exec.Cmd, string, time.Duration) error { return errWindowsOnly }
func signalReady() error                                       { return nil }
