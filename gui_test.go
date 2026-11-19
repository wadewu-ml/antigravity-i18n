package main

import (
	"errors"
	"testing"

	"antigravity-i18n/internal/core"
)

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
