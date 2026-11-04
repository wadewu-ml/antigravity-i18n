//go:build !windows

package core

import "os/exec"

// hideWindow is a no-op outside Windows, where tasklist and taskkill are not used.
func hideWindow(cmd *exec.Cmd) {}
