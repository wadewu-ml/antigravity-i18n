package selfupdate

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func put(t *testing.T, path, data string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(data), 0600); err != nil {
		t.Fatal(err)
	}
}

func contents(t *testing.T, path string) string {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}

func TestSwapAndRecovery(t *testing.T) {
	for _, scenario := range []string{"success", "stage missing", "launch fails"} {
		t.Run(scenario, func(t *testing.T) {
			exe := filepath.Join(t.TempDir(), "我的工具.exe")
			personal := exe + ".old.exe"
			put(t, exe, "old")
			put(t, personal, "user copy")
			tx, err := newReceipt(exe, "4.0.3", "v4.1.0")
			if err != nil {
				t.Fatal(err)
			}
			if scenario != "stage missing" {
				put(t, tx.staged(exe), "new")
			}
			started := false
			err = replaceAndStart(exe, tx, func() error {
				started = true
				if contents(t, exe) != "new" || contents(t, tx.backup(exe)) != "old" {
					t.Fatal("incorrect handoff")
				}
				if scenario == "launch fails" {
					return errors.New("launch failed")
				}
				return nil
			})
			if started != (scenario != "stage missing") {
				t.Fatal("unexpected launch")
			}
			if scenario == "success" {
				if err != nil {
					t.Fatal(err)
				}
				if err = cleanup(exe, "4.0.3"); err == nil {
					t.Fatal("old GUI removed its own backup")
				}
				if err = cleanup(exe, "4.1.0"); err != nil {
					t.Fatal(err)
				}
			} else {
				if err == nil || contents(t, exe) != "old" {
					t.Fatalf("rollback: %v", err)
				}
				if err = cleanup(exe, "4.0.3"); err != nil {
					t.Fatal(err)
				}
			}
			entries, err := os.ReadDir(filepath.Dir(exe))
			if err != nil || len(entries) != 2 || contents(t, personal) != "user copy" {
				t.Fatalf("cleanup: %v %v", entries, err)
			}
		})
	}
}

func TestConcurrentUpdate(t *testing.T) {
	exe := filepath.Join(t.TempDir(), "tool.exe")
	unlock, err := lockExecutable(exe)
	if err != nil {
		t.Fatal(err)
	}
	if release, err := lockExecutable(exe); err == nil {
		release()
		unlock()
		t.Fatal("locked twice")
	}
	unlock()
	release, err := lockExecutable(exe)
	if err != nil {
		t.Fatal(err)
	}
	release()
}

func TestReceiptPublicationPreservesExistingRecord(t *testing.T) {
	dir := t.TempDir()
	exe := filepath.Join(dir, "tool.exe")
	put(t, receiptPath(exe), "existing record")
	if _, err := newReceipt(exe, "4.0.3", "v4.1.0"); err == nil {
		t.Fatal("published a second receipt")
	}
	if got := contents(t, receiptPath(exe)); got != "existing record" {
		t.Fatalf("changed the existing receipt: %q", got)
	}
	if entries, err := os.ReadDir(dir); err != nil || len(entries) != 1 {
		t.Fatalf("unexpected files after failed publication: %v, %v", entries, err)
	}
}

func TestRollbackPreservesOccupiedOriginalName(t *testing.T) {
	exe := filepath.Join(t.TempDir(), "tool.exe")
	tx, err := newReceipt(exe, "4.0.3", "v4.1.0")
	if err != nil {
		t.Fatal(err)
	}
	put(t, exe, "old")
	put(t, tx.staged(exe), "new")
	err = replaceAndStart(exe, tx, func() error {
		put(t, tx.staged(exe), "occupied")
		return errors.New("launch failed")
	})
	if err == nil || contents(t, tx.backup(exe)) != "old" || contents(t, tx.staged(exe)) != "occupied" {
		t.Fatalf("recovery backup: %v", err)
	}
}
