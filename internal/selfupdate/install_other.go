//go:build !windows

package selfupdate

import (
	"errors"
	"os/exec"
)

var errWindowsOnly = errors.New("工具自更新仅支持 Windows。")

func lockExecutable(string) (func(), error) { return nil, errWindowsOnly }
func hideProbeWindow(*exec.Cmd)             {}
func moveNewName(string, string) error      { return errWindowsOnly }
