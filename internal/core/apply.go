package core

import (
	"fmt"
	"os"
	"path/filepath"
)

// ApplyLocale installs a locale into Antigravity, replacing whatever patch is
// present.
func (s *Service) ApplyLocale(opts Options) (*ApplyResult, error) {
	s.clearNotes()

	p, err := resolveAppPaths(opts.AppDir)
	if err != nil {
		return nil, err
	}

	localeCode := opts.Locale
	if localeCode == "" {
		localeCode = DefaultLocale
	}
	locale, err := s.store.LoadLocale(localeCode)
	if err != nil {
		return nil, err
	}

	// Fragments are built before anything is modified, so a malformed locale
	// fails fast instead of producing a broken UI after the rewrite.
	preloadFragment, err := s.store.buildPreloadFragment(locale)
	if err != nil {
		return nil, err
	}
	menuFragment, err := s.store.buildMenuFragment(locale)
	if err != nil {
		return nil, err
	}
	dialogFragment, err := s.store.buildDialogFragment(locale)
	if err != nil {
		return nil, err
	}

	// The generated JavaScript goes into the app's own scripts, so it is checked
	// here rather than being discovered by a failed launch.
	for _, f := range []struct{ name, src string }{
		{"dist/preload.js", preloadFragment},
		{"dist/menu.js", menuFragment},
		{"dist/main.js", dialogFragment},
	} {
		if err := checkJavaScriptSyntax(f.name, f.src); err != nil {
			return nil, err
		}
	}

	initialHash, err := hashArchive(p.AsarPath)
	if err != nil {
		return nil, err
	}
	if detectPatchState(p.AsarPath) == nil {
		return nil, fmt.Errorf("could not read Antigravity archive: %s", p.AsarPath)
	}
	expectedIdentity, err := readArchiveIdentity(p.AsarPath)
	if err != nil {
		return nil, fmt.Errorf("could not verify Antigravity archive identity: %s", p.AsarPath)
	}
	if expectedIdentity.Name != "antigravity" {
		return nil, fmt.Errorf("could not verify Antigravity archive identity: %s", p.AsarPath)
	}

	timeout := opts.TimeoutMs
	if timeout <= 0 {
		timeout = gracefulTimeoutMs
	}
	wasRunning, stopped, forced, err := stopAntigravity(opts.Force, timeout)
	if err != nil {
		return nil, err
	}
	if !stopped {
		return nil, fmt.Errorf("Antigravity is still running; close it and re-run")
	}
	if err := assertArchiveUnchanged(p.AsarPath, initialHash); err != nil {
		return nil, err
	}

	currentPatched := detectPatchState(p.AsarPath)
	if currentPatched == nil {
		return nil, fmt.Errorf("could not read Antigravity archive after shutdown: %s", p.AsarPath)
	}

	txn := newFileTxn()
	for _, path := range []string{
		p.AsarPath,
		p.CleanBackupPath,
		p.StatePath,
	} {
		if err := txn.track(path); err != nil {
			return nil, err
		}
	}

	// The clean backup is what restore needs, so it has to be a verified
	// unpatched archive from the current release before anything is rewritten.
	if err := ensureCleanBackupForPatch(s, p, p.CleanBackupPath, *currentPatched, expectedIdentity, txn); err != nil {
		return nil, txn.failWithRollback("preparing the clean backup", err)
	}

	backupPath := createUniqueBackupPath(p.ResourcesDir, fmt.Sprintf("app.asar.bak-%s", formatTimestamp()))
	if err := copyFileExclusive(p.AsarPath, backupPath); err != nil {
		return nil, txn.failWithRollback("creating a safety backup", err)
	}
	noteBackupAccumulation(s, p.ResourcesDir)

	staged, err := buildPatchedArchive(p.AsarPath, locale, preloadFragment, menuFragment, dialogFragment)
	if err != nil {
		return nil, txn.failWithRollback("building the patched archive", err)
	}
	defer os.Remove(staged)

	// Verify the staged archive in full before it replaces the live one.
	arc, err := ReadArchive(staged)
	if err != nil {
		return nil, txn.failWithRollback("verifying the patched archive", err)
	}
	if detectPatchState(staged) == nil || !*detectPatchState(staged) {
		return nil, txn.failWithRollback("verifying the patched archive",
			fmt.Errorf("localization markers were not found in the patched archive"))
	}
	if !SameRelease(expectedIdentity, arc.Identity()) {
		return nil, txn.failWithRollback("verifying the patched archive",
			fmt.Errorf("application identity changed in the patched archive"))
	}
	if _, _, ok := readArchiveLocale(staged); !ok {
		return nil, txn.failWithRollback("verifying the patched archive",
			fmt.Errorf("the patched archive does not declare a locale"))
	}
	for _, rel := range []string{"dist/main.js", "dist/menu.js", "dist/preload.js"} {
		data, err := arc.ReadFile(rel)
		if err != nil {
			return nil, txn.failWithRollback("verifying the patched archive", err)
		}
		if err := checkJavaScriptSyntax(rel, string(data)); err != nil {
			return nil, txn.failWithRollback("verifying the patched archive", err)
		}
	}

	packedHash, err := hashArchive(staged)
	if err != nil {
		return nil, txn.failWithRollback("verifying the patched archive", err)
	}

	if err := atomicReplaceFile(staged, p.AsarPath, func(candidate string) error {
		running, err := isAntigravityRunning()
		if err != nil {
			return err
		}
		if running {
			return fmt.Errorf("Antigravity restarted during this operation. Close it and re-run")
		}
		if err := assertArchiveUnchanged(candidate, packedHash); err != nil {
			return err
		}
		return assertArchiveUnchanged(p.AsarPath, initialHash)
	}); err != nil {
		return nil, txn.failWithRollback("installing the patched archive", err)
	}

	writeState(p, locale.Language)

	result := &ApplyResult{
		Locale:      locale.Language,
		LocaleName:  locale.Name,
		BackupPath:  backupPath,
		CleanBackup: p.CleanBackupPath,
		WasRunning:  wasRunning,
		Forced:      forced,
	}

	if opts.Restart {
		launchAntigravity(p)
	}
	return result, nil
}

// readArchiveIdentity reads the package.json identity out of an archive.
func readArchiveIdentity(asarPath string) (*Identity, error) {
	arc, err := ReadArchive(asarPath)
	if err != nil {
		return nil, err
	}
	id := arc.Identity()
	if id == nil {
		return nil, fmt.Errorf("archive has no readable identity")
	}
	return id, nil
}

// buildPatchedArchive produces a patched copy of the archive in a temporary
// directory and returns its path.
func buildPatchedArchive(srcPath string, locale *Locale, preloadFragment, menuFragment, dialogFragment string) (string, error) {
	arc, err := ReadArchive(srcPath)
	if err != nil {
		return "", err
	}

	main, err := arc.ReadFile("dist/main.js")
	if err != nil {
		return "", err
	}
	menu, err := arc.ReadFile("dist/menu.js")
	if err != nil {
		return "", err
	}
	preload, err := arc.ReadFile("dist/preload.js")
	if err != nil {
		return "", err
	}

	newMain, err := patchMain(string(main), locale, dialogFragment)
	if err != nil {
		return "", err
	}
	newMenu, err := patchMenu(string(menu), menuFragment)
	if err != nil {
		return "", err
	}
	newPreload := patchPreload(string(preload), preloadFragment)

	tmp, err := os.CreateTemp("", "antigravity-i18n-patched-*.asar")
	if err != nil {
		return "", err
	}
	staged := tmp.Name()
	if err := tmp.Close(); err != nil {
		os.Remove(staged)
		return "", err
	}

	// Every body except these three is copied from the source unchanged, so the
	// rebuilt archive stays byte identical apart from the patch itself.
	err = WritePatched(srcPath, staged, []Patch{
		{Path: "dist/main.js", Data: []byte(newMain)},
		{Path: "dist/menu.js", Data: []byte(newMenu)},
		{Path: "dist/preload.js", Data: []byte(newPreload)},
	})
	if err != nil {
		os.Remove(staged)
		return "", err
	}
	return staged, nil
}

// ensureCleanBackupForPatch makes sure a verified unpatched archive from the
// current release exists at cleanBackupPath.
func ensureCleanBackupForPatch(s *Service, p *paths, cleanBackupPath string, currentPatched bool, expectedIdentity *Identity, txn *fileTxn) error {
	if !currentPatched {
		// An official update can leave an older clean backup beside a new
		// app.asar. Refreshing from the verified-unpatched archive keeps a later
		// restore on the same Antigravity release.
		if err := copyFileExclusive(p.AsarPath, cleanBackupPath); err != nil {
			// A backup that is already valid for this release is fine to keep.
			if fileExists(cleanBackupPath) && inspectCleanBackup(cleanBackupPath, expectedIdentity) == cleanBackupValid {
				return nil
			}
			return err
		}
		return nil
	}

	if inspectCleanBackup(cleanBackupPath, expectedIdentity) == cleanBackupValid {
		return nil
	}

	pristine, err := findPristineArchive(p.ResourcesDir, expectedIdentity)
	if err != nil {
		return err
	}
	if pristine == "" {
		return fmt.Errorf(
			"app.asar is already patched and no verified pristine copy was found, so the original " +
				"English build cannot be preserved. Reinstall or update Antigravity to restore a clean " +
				"app.asar, then run this tool again")
	}
	if err := copyFileExclusive(pristine, cleanBackupPath); err != nil {
		return err
	}
	s.notes = append(s.notes, fmt.Sprintf("Recovered a pristine archive from %s", filepath.Base(pristine)))
	return nil
}

// RestoreOfficial restores the official untranslated archive.
func (s *Service) RestoreOfficial(opts Options) (*RestoreResult, error) {
	s.clearNotes()

	p, err := resolveAppPaths(opts.AppDir)
	if err != nil {
		return nil, err
	}

	initialHash, err := hashArchive(p.AsarPath)
	if err != nil {
		return nil, err
	}
	currentPatched := detectPatchState(p.AsarPath)
	if currentPatched == nil {
		return nil, fmt.Errorf("could not read Antigravity archive: %s", p.AsarPath)
	}

	// Verify the archive before refreshing or restoring a clean backup. Without
	// this guard, a mistaken --app-dir pointing at an unrelated app.asar could
	// overwrite a known-good backup even though the live archive is untouched.
	expectedIdentity, err := readArchiveIdentity(p.AsarPath)
	if err != nil {
		return nil, fmt.Errorf("could not verify Antigravity archive identity: %s", p.AsarPath)
	}
	if expectedIdentity.Name != "antigravity" {
		return nil, fmt.Errorf("could not verify Antigravity archive identity: %s", p.AsarPath)
	}

	if !*currentPatched {
		// After an official update the app is already English, while the clean
		// backup may still belong to the previous release. Refresh the backup and
		// avoid needlessly stopping or rewriting the running application.
		if err := ensureRestorableCleanBackup(s, p, expectedIdentity); err != nil {
			return nil, err
		}
		writeState(p, "en")
		return &RestoreResult{AlreadyEnglish: true}, nil
	}

	timeout := opts.TimeoutMs
	if timeout <= 0 {
		timeout = gracefulTimeoutMs
	}

	txn := newFileTxn()
	for _, path := range []string{p.AsarPath, p.CleanBackupPath, p.StatePath} {
		if err := txn.track(path); err != nil {
			return nil, err
		}
	}

	if err := ensureRestorableCleanBackup(s, p, expectedIdentity); err != nil {
		return nil, txn.failWithRollback("locating a clean English archive", err)
	}

	backupHash, err := hashArchive(p.CleanBackupPath)
	if err != nil {
		return nil, txn.failWithRollback("reading the clean backup", err)
	}

	wasRunning, stopped, forced, err := stopAntigravity(opts.Force, timeout)
	if err != nil {
		return nil, err
	}
	if !stopped {
		return nil, fmt.Errorf("Antigravity is still running; close it and re-run")
	}
	if err := assertArchiveUnchanged(p.AsarPath, initialHash); err != nil {
		return nil, err
	}

	backupPath := createUniqueBackupPath(p.ResourcesDir,
		fmt.Sprintf("app.asar.bak-before-restore-%s", formatTimestamp()))
	if err := copyFileExclusive(p.AsarPath, backupPath); err != nil {
		return nil, txn.failWithRollback("creating a safety backup", err)
	}
	noteBackupAccumulation(s, p.ResourcesDir)

	if err := atomicReplaceFile(p.CleanBackupPath, p.AsarPath, func(candidate string) error {
		running, err := isAntigravityRunning()
		if err != nil {
			return err
		}
		if running {
			return fmt.Errorf("Antigravity restarted during this operation. Close it and re-run")
		}
		if err := assertArchiveUnchanged(candidate, backupHash); err != nil {
			return err
		}
		return assertArchiveUnchanged(p.AsarPath, initialHash)
	}); err != nil {
		return nil, txn.failWithRollback("restoring the clean archive", err)
	}

	writeState(p, "en")

	result := &RestoreResult{
		BackupPath: backupPath,
		WasRunning: wasRunning,
		Forced:     forced,
	}
	if opts.Restart {
		launchAntigravity(p)
	}
	return result, nil
}

// ensureRestorableCleanBackup makes sure a verified pristine archive is
// available at the clean backup path.
func ensureRestorableCleanBackup(s *Service, p *paths, expectedIdentity *Identity) error {
	if inspectCleanBackup(p.CleanBackupPath, expectedIdentity) == cleanBackupValid {
		return nil
	}
	pristine, err := findPristineArchive(p.ResourcesDir, expectedIdentity)
	if err != nil {
		return err
	}
	if pristine == "" {
		return fmt.Errorf(
			"no verified pristine English app.asar was found. Reinstall or update Antigravity to " +
				"restore the official build, then run this tool again")
	}
	if err := copyFileExclusive(pristine, p.CleanBackupPath); err != nil && !fileExists(p.CleanBackupPath) {
		return err
	}
	s.notes = append(s.notes, fmt.Sprintf("Using verified original backup: %s", filepath.Base(pristine)))
	return nil
}
