package main

import (
	"errors"
	"testing"

	"antigravity-i18n/internal/core"
)

func TestFormatAppErrorForWails(t *testing.T) {
	timeout := &core.OpError{Op: "closing Antigravity", Cause: core.ErrGracefulShutdownTimeout}
	if got := formatAppError(timeout); got != "shutdown_timeout: "+timeout.Error() {
		t.Fatalf("shutdown timeout must travel as tagged text, got %#v", got)
	}

	rollback := &core.RollbackError{
		Op:           "restoring",
		Cause:        timeout,
		RollbackErrs: []error{errors.New("rollback failed")},
	}
	if got := formatAppError(rollback); got != rollback.Error() {
		t.Fatalf("rollback errors should report their details, got %#v", got)
	}
}
