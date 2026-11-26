//go:build windows

package core

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"golang.org/x/sys/windows"
)

const archiveHolderEnv = "REGRAVITY_TEST_ARCHIVE_HOLDER"

func TestMain(m *testing.M) {
	if os.Getenv(archiveHolderEnv) != "" {
		// The helper can start with the Sidecar SDK's arguments in its Windows
		// command line. Only the Go test runner needs to see the test flag.
		os.Args = []string{os.Args[0], "-test.run=^TestArchiveHolderProcess$"}
	}
	os.Exit(m.Run())
}

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

func startArchiveHolder(t *testing.T, p *paths, path string, args ...string) uint32 {
	t.Helper()
	self, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}
	fixtureExe := filepath.Join(p.AppDir, antigravityExe)
	copyFixture(t, self, fixtureExe)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	t.Cleanup(cancel)
	holder := exec.CommandContext(ctx, fixtureExe, args...)
	holder.Env = append(os.Environ(), archiveHolderEnv+"="+path)
	hideWindow(holder)
	stdout, err := holder.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}
	if err := holder.Start(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		_ = holder.Process.Kill()
		_ = holder.Wait()
	})
	scanner := bufio.NewScanner(stdout)
	if !scanner.Scan() || scanner.Text() != "archive open" {
		t.Fatalf("isolated process did not open its archive: %q, %v", scanner.Text(), scanner.Err())
	}
	return uint32(holder.Process.Pid)
}

func TestForceApplyClosesArchiveHolder(t *testing.T) {
	p := fixtureInstall(t, syntheticAsar)
	startArchiveHolder(t, p, p.AsarPath)
	svc := newTestService(t)
	svc.appRunning = isAntigravityRunning
	svc.stopApp = func(*paths, bool, int) (bool, error) {
		t.Fatal("require-closed apply must only check for a running process")
		return false, nil
	}
	opts := Options{AppDir: p.AppDir, Locale: EnglishLocale, SessionMetrics: new(true), RequireClosed: true, TimeoutMs: 500}
	if _, err := svc.ApplyLocale(opts); !errors.Is(err, ErrAntigravityRunning) {
		t.Fatalf("expected an immediate running-app error, got %v", err)
	}
	svc.stopApp = stopAntigravity
	opts.RequireClosed = false
	if _, err := svc.ApplyLocale(opts); !errors.Is(err, ErrGracefulShutdownTimeout) {
		t.Fatalf("expected a normal shutdown timeout, got %v", err)
	}
	opts.Force = true
	opts.RequireClosed = true
	opts.TimeoutMs = 5000
	applied, err := svc.ApplyLocale(opts)
	if err != nil || !applied.Forced {
		t.Fatalf("force apply with a locked archive: %+v, %v", applied, err)
	}
}

func TestApplyWithAntigravitySidecar(t *testing.T) {
	p := fixtureInstall(t, syntheticAsar)
	dataPath := filepath.Join(p.AppDir, "sidecar-data")
	if err := os.WriteFile(dataPath, nil, 0o600); err != nil {
		t.Fatal(err)
	}
	pid := startArchiveHolder(t, p, dataPath,
		"--import=file:///C:/Users/Test/.gemini/antigravity/lib/node_modules/sidecar_sdk/register_loader.js",
		"main.mjs")

	svc := newTestService(t)
	svc.stopApp = stopAntigravity
	svc.appRunning = isAntigravityRunning
	applied, err := svc.ApplyLocale(Options{AppDir: p.AppDir, Locale: EnglishLocale, SessionMetrics: new(true), TimeoutMs: 2000})
	if err != nil || applied.WasRunning {
		t.Fatalf("apply with a running Sidecar SDK worker: %+v, %v", applied, err)
	}
	if processIDExited(pid) {
		t.Fatal("applying settings stopped the Sidecar SDK worker")
	}
}
