package core

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func newTestService(t *testing.T) *Service {
	t.Helper()
	assets := os.DirFS("../..")
	store, err := NewLocaleStore(assets, assets, "assets/locales", "assets/patches")
	if err != nil {
		t.Fatal(err)
	}
	svc := NewService(store)
	svc.stopApp = func(*paths, bool, int) (bool, error) { return false, nil }
	svc.appRunning = func(*paths) (bool, error) { return false, nil }
	return svc
}

const syntheticAsar = "../../testdata/synthetic-app.asar"

func fixtureInstall(t *testing.T, archive string) *paths {
	t.Helper()
	appDir := t.TempDir()
	resources := filepath.Join(appDir, "resources")
	if err := os.Mkdir(resources, 0o755); err != nil {
		t.Fatal(err)
	}
	copyFixture(t, archive, filepath.Join(resources, "app.asar"))
	p, err := resolveAppPaths(appDir)
	if err != nil {
		t.Fatal(err)
	}
	return p
}

func copyFixture(t *testing.T, src, dest string) {
	t.Helper()
	data, err := os.ReadFile(src)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(dest, data, 0o644); err != nil {
		t.Fatal(err)
	}
}

func fileHash(t *testing.T, path string) string {
	t.Helper()
	sum, err := hashArchive(path)
	if err != nil {
		t.Fatal(err)
	}
	return sum
}

func TestApplySwitchAndRestore(t *testing.T) {
	for _, example := range []struct{ name, archive string }{
		{"synthetic", syntheticAsar},
		{"official", os.Getenv("ANTIGRAVITY_ASAR")},
	} {
		t.Run(example.name, func(t *testing.T) {
			if example.archive == "" {
				t.Skip("set ANTIGRAVITY_ASAR to a clean official archive")
			}
			p := fixtureInstall(t, example.archive)
			original, err := ReadArchive(p.AsarPath)
			if err != nil {
				t.Fatal(err)
			}
			identity := original.Identity()
			if identity == nil {
				t.Fatal("fixture has no identity")
			}
			main, err := original.ReadFile("dist/main.js")
			if err != nil {
				t.Fatal(err)
			}
			// A clean backup with the same version can still be a different build.
			stale := append(main, []byte("\n// earlier build of the same version\n")...)
			if err := WritePatched(p.AsarPath, p.CleanBackupPath, []Patch{{Path: "dist/main.js", Data: stale}}); err != nil {
				t.Fatal(err)
			}

			svc := newTestService(t)
			opts := Options{AppDir: p.AppDir, Locale: DefaultLocale}
			want := fileHash(t, p.AsarPath)
			if _, err := svc.ApplyLocale(opts); err != nil {
				t.Fatal(err)
			}
			if fileHash(t, p.CleanBackupPath) != want {
				t.Fatal("the same-version backup was not refreshed from the installed archive")
			}
			opts.Locale = "ja"
			opts.SessionMetrics = new(true)
			if _, err := svc.ApplyLocale(opts); err != nil {
				t.Fatal(err)
			}
			status, err := svc.GetStatus(p.AppDir)
			if err != nil || status.CurrentLanguage != "ja" || status.AppVersion != identity.Version || !status.SessionMetrics {
				t.Fatalf("status after switching locale: %+v, %v", status, err)
			}
			patched, err := ReadArchive(p.AsarPath)
			if err != nil {
				t.Fatal(err)
			}
			for _, rel := range []string{"dist/preload.js", "dist/ideInstall/wizardPreload.js", "dist/remote/sshPromptPreload.js"} {
				if original.lookup(rel) == nil {
					continue
				}
				preload, err := patched.ReadFile(rel)
				if err != nil || strings.Count(string(preload), "/* antigravity-i18n:begin */") != 1 {
					t.Fatalf("%s does not contain exactly one engine after switching: %v", rel, err)
				}
				if strings.Contains(string(preload), sessionMetricsMarker) != (rel == "dist/preload.js") {
					t.Fatalf("unexpected session metrics state in %s", rel)
				}
			}
			if original.lookup("dist/remote/sshPromptPreload.js") != nil {
				prompt, err := patched.ReadFile("dist/remote/sshPrompts.js")
				if err != nil || !strings.Contains(string(prompt), `<div id="message" translate="no"></div>`) {
					t.Fatalf("SSH prompt content is not protected: %v", err)
				}
			}
			first := fileHash(t, p.AsarPath)
			opts.SessionMetrics = nil
			if _, err := svc.ApplyLocale(opts); err != nil || fileHash(t, p.AsarPath) != first {
				t.Fatalf("re-applying the locale was not byte-stable: %v", err)
			}
			opts.Locale = EnglishLocale
			opts.SessionMetrics = new(true)
			if _, err := svc.ApplyLocale(opts); err != nil {
				t.Fatal(err)
			}
			status, err = svc.GetStatus(p.AppDir)
			if err != nil || status.CurrentLanguage != EnglishLocale || !status.SessionMetrics {
				t.Fatalf("status with metrics alone: %+v, %v", status, err)
			}
			english, err := ReadArchive(p.AsarPath)
			if err != nil {
				t.Fatal(err)
			}
			clean, err := ReadArchive(p.CleanBackupPath)
			if err != nil {
				t.Fatal(err)
			}
			for _, rel := range []string{"dist/menu.js", "dist/ideInstall/wizardPreload.js", "dist/remote/sshPromptPreload.js", "dist/remote/sshPrompts.js"} {
				if clean.lookup(rel) == nil {
					continue
				}
				pristine, err := clean.ReadFile(rel)
				if err != nil {
					t.Fatal(err)
				}
				actual, err := english.ReadFile(rel)
				if err != nil || !bytes.Equal(pristine, actual) {
					t.Fatalf("metrics-only apply left a language patch in %s: %v", rel, err)
				}
			}
			opts.Locale = "ja"
			opts.SessionMetrics = new(false)
			if _, err := svc.ApplyLocale(opts); err != nil {
				t.Fatal(err)
			}
			status, err = svc.GetStatus(p.AppDir)
			if err != nil || status.SessionMetrics || status.CurrentLanguage != "ja" {
				t.Fatalf("status after disabling metrics: %+v, %v", status, err)
			}
			if result, err := svc.RestoreOfficial(opts); err != nil || result.AlreadyEnglish || fileHash(t, p.AsarPath) != want {
				t.Fatalf("restoring Antigravity %s did not reproduce the official bytes: %v", identity.Version, err)
			}
			t.Logf("Antigravity %s restored byte-exactly", identity.Version)
		})
	}
}

func TestRestoreRecoversPristineBackup(t *testing.T) {
	p := fixtureInstall(t, syntheticAsar)
	want := fileHash(t, p.AsarPath)
	svc := newTestService(t)
	opts := Options{AppDir: p.AppDir}
	result, err := svc.ApplyLocale(opts)
	if err != nil {
		t.Fatal(err)
	}
	copyFixture(t, p.AsarPath, p.CleanBackupPath)
	writeOtherRelease(t, filepath.Join(p.ResourcesDir, "app.asar.bak-99991231-235959"))

	hidden := filepath.Join(p.ResourcesDir, "saved-original.asar")
	if err := os.Rename(result.BackupPath, hidden); err != nil {
		t.Fatal(err)
	}
	localized := fileHash(t, p.AsarPath)
	if _, err := svc.RestoreOfficial(opts); err == nil || fileHash(t, p.AsarPath) != localized {
		t.Fatal("restore used a patched backup or a different release")
	}
	if err := os.Rename(hidden, result.BackupPath); err != nil {
		t.Fatal(err)
	}
	if _, err := svc.RestoreOfficial(opts); err != nil || fileHash(t, p.AsarPath) != want {
		t.Fatalf("restore failed to recover the original timestamped backup: %v", err)
	}
}

func writeOtherRelease(t *testing.T, dst string) {
	t.Helper()
	pkg := []byte(`{"name":"antigravity","version":"9.9.8"}`)
	if err := WritePatched(syntheticAsar, dst, []Patch{{Path: "package.json", Data: pkg}}); err != nil {
		t.Fatal(err)
	}
}

func TestOverlappingOperationsLeaveBackupIntact(t *testing.T) {
	p := fixtureInstall(t, syntheticAsar)
	want := fileHash(t, p.AsarPath)
	first, second := newTestService(t), newTestService(t)
	opts := Options{AppDir: p.AppDir}
	overlapped := false
	first.Progress = func(stage, _ string) {
		if stage != "正在写回 app.asar" {
			return
		}
		overlapped = true
		if _, err := second.ApplyLocale(Options{AppDir: p.AppDir, Locale: "ja"}); !errors.Is(err, errInstallBusy) {
			t.Fatalf("overlapping apply: %v", err)
		}
		if _, err := second.RestoreOfficial(opts); !errors.Is(err, errInstallBusy) {
			t.Fatalf("overlapping restore: %v", err)
		}
	}
	if _, err := first.ApplyLocale(opts); err != nil || !overlapped {
		t.Fatalf("could not exercise overlapping operations: %v", err)
	}
	if fileHash(t, p.CleanBackupPath) != want {
		t.Fatal("a rejected operation changed the clean backup")
	}
	if _, err := os.Stat(p.StatePath); err != nil {
		t.Fatal(err)
	}
	if _, err := second.RestoreOfficial(opts); err != nil || fileHash(t, p.AsarPath) != want {
		t.Fatalf("restore after the lock was released: %v", err)
	}
}

func TestOperationsPreserveConcurrentArchiveUpdate(t *testing.T) {
	for _, action := range []string{"apply", "restore"} {
		t.Run(action, func(t *testing.T) {
			p := fixtureInstall(t, syntheticAsar)
			svc := newTestService(t)
			opts := Options{AppDir: p.AppDir}
			if action == "restore" {
				if _, err := svc.ApplyLocale(opts); err != nil {
					t.Fatal(err)
				}
			}
			updated := filepath.Join(t.TempDir(), "updated.asar")
			writeOtherRelease(t, updated)
			svc.Progress = func(stage, _ string) {
				if stage == "正在写回 app.asar" {
					copyFixture(t, updated, p.AsarPath)
				}
			}
			var err error
			if action == "apply" {
				_, err = svc.ApplyLocale(opts)
			} else {
				_, err = svc.RestoreOfficial(opts)
			}
			if err == nil || fileHash(t, p.AsarPath) != fileHash(t, updated) {
				t.Fatalf("external archive replacement was not preserved: %v", err)
			}
		})
	}
}

func TestStatusRejectsInvalidArchive(t *testing.T) {
	svc := newTestService(t)
	t.Run("truncated archive", func(t *testing.T) {
		p := fixtureInstall(t, syntheticAsar)
		if err := os.Truncate(p.AsarPath, 10); err != nil {
			t.Fatal(err)
		}
		if status, err := svc.GetStatus(p.AppDir); err == nil {
			t.Fatalf("reported status for a truncated archive: %+v", status)
		}
	})
	t.Run("unrelated application", func(t *testing.T) {
		p := fixtureInstall(t, syntheticAsar)
		pkg := []byte(`{"name":"other-app","version":"1.0.0"}`)
		if err := WritePatched(syntheticAsar, p.AsarPath, []Patch{{Path: "package.json", Data: pkg}}); err != nil {
			t.Fatal(err)
		}
		if status, err := svc.GetStatus(p.AppDir); err == nil {
			t.Fatalf("reported status for another application: %+v", status)
		}
	})
}
