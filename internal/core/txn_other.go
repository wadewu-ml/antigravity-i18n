//go:build !windows

package core

import (
	"errors"
	"os"

	"golang.org/x/sys/unix"
)

func tryLockFile(file *os.File) (func(), error) {
	fd := int(file.Fd())
	err := unix.Flock(fd, unix.LOCK_EX|unix.LOCK_NB)
	if errors.Is(err, unix.EWOULDBLOCK) {
		return nil, errInstallBusy
	}
	if err != nil {
		return nil, err
	}
	return func() {
		_ = unix.Flock(fd, unix.LOCK_UN)
	}, nil
}
