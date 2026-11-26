package main

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"regravity/internal/core"
)

func TestInitialDetectionKeepsChosenInstallation(t *testing.T) {
	archive, err := os.ReadFile("testdata/synthetic-app.asar")
	if err != nil {
		t.Fatal(err)
	}
	detected, selected := t.TempDir(), t.TempDir()
	for _, dir := range []string{detected, selected} {
		if err := os.WriteFile(filepath.Join(dir, "app.asar"), archive, 0o600); err != nil {
			t.Fatal(err)
		}
	}
	t.Setenv("ANTIGRAVITY_APP_DIR", detected)
	app := NewApp(core.NewService(nil), "")
	if _, err := app.GetStatus(); err != nil || app.currentDir() != detected {
		t.Fatalf("initial detection selected %q: %v", app.currentDir(), err)
	}
	if _, err := app.SetAppDir(selected); err != nil {
		t.Fatal(err)
	}
	// Finish an automatic detection after a manual selection has committed.
	status, err := app.detectInitialStatus()
	if err != nil || status.AppDir != detected {
		t.Fatalf("delayed detection returned %+v: %v", status, err)
	}
	if got := app.operationOptions(false).AppDir; got != selected {
		t.Fatalf("operation would target %q instead of %q", got, selected)
	}
}

func TestFormatAppErrorForWails(t *testing.T) {
	running := core.ErrAntigravityRunning
	if got := formatAppError(running); got != "antigravity_running: "+running.Error() {
		t.Fatalf("running app must travel as tagged text, got %#v", got)
	}

	rollback := &core.RollbackError{
		Op:           "restoring",
		Cause:        running,
		RollbackErrs: []error{errors.New("rollback failed")},
	}
	if got := formatAppError(rollback); got != rollback.Error() {
		t.Fatalf("rollback errors should report their details, got %#v", got)
	}
}
