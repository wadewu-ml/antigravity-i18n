//go:build windows

package core

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"golang.org/x/sys/windows"
)

const archiveHolderEnv = "ANTIGRAVITY_I18N_TEST_ARCHIVE_HOLDER"

func TestArchiveHolderProcess(t *testing.T) {
	path := os.Getenv(archiveHolderEnv)
	if path == "" {
		return
	}
	name, err := windows.UTF16PtrFromString(path)
	if err != nil {
		t.Fatal(err)
	}
	handle, err := windows.CreateFile(name, windows.GENERIC_READ, windows.FILE_SHARE_READ, nil, windows.OPEN_EXISTING, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer windows.CloseHandle(handle)
	fmt.Println("archive open")
	time.Sleep(30 * time.Second)
}

func TestForceApplyClosesArchiveHolder(t *testing.T) {
	p := fixtureInstall(t, syntheticAsar)
	self, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}
	fixtureExe := filepath.Join(p.AppDir, antigravityExe)
	copyFixture(t, self, fixtureExe)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	holder := exec.CommandContext(ctx, fixtureExe, "-test.run=^TestArchiveHolderProcess$")
	holder.Env = append(os.Environ(), archiveHolderEnv+"="+p.AsarPath)
	hideWindow(holder)
	stdout, err := holder.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}
	if err := holder.Start(); err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = holder.Process.Kill()
		_ = holder.Wait()
	}()
	scanner := bufio.NewScanner(stdout)
	if !scanner.Scan() || scanner.Text() != "archive open" {
		t.Fatalf("isolated process did not open its archive: %q, %v", scanner.Text(), scanner.Err())
	}

	svc := newTestService(t)
	svc.stopApp = stopAntigravity
	svc.appRunning = isAntigravityRunning
	applied, err := svc.ApplyLocale(Options{AppDir: p.AppDir, Locale: EnglishLocale, Force: true, SessionMetrics: new(true), TimeoutMs: 5000})
	if err != nil || !applied.Forced {
		t.Fatalf("force apply with a locked archive: %+v, %v", applied, err)
	}
}
