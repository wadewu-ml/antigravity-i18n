package core

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ApplyLocale installs the selected language and optional features.
func (s *Service) ApplyLocale(opts Options) (*ApplyResult, error) {
	s.clearNotes()

	p, err := resolveAppPaths(opts.AppDir)
	if err != nil {
		return nil, err
	}
	unlock, err := lockInstallation(p.ResourcesDir)
	if err != nil {
		return nil, err
	}
	defer unlock()

	localeCode := opts.Locale
	if localeCode == "" {
		localeCode = DefaultLocale
	}
	initialHash, err := hashArchive(p.AsarPath)
	if err != nil {
		return nil, err
	}
	expectedIdentity, err := readArchiveIdentity(p.AsarPath)
	if err != nil {
		return nil, err
	}
	currentPatched, err := detectPatchInArchive(p.AsarPath)
	if err != nil {
		return nil, err
	}
	var features archiveFeatures
	if currentPatched && (opts.SessionMetrics == nil || opts.EffortSlider == nil) {
		_, _, features, _ = readArchiveLocale(p.AsarPath)
	}
	if opts.SessionMetrics != nil {
		features.sessionMetrics = *opts.SessionMetrics
	}
	if opts.EffortSlider != nil {
		features.effortSlider = *opts.EffortSlider
	}
	english := CanonicalLocale(localeCode) == EnglishLocale
	var locale *Locale
	var patch patchFragments
	if english {
		if features == (archiveFeatures{}) {
			restored, err := s.restoreOfficial(p, opts, currentPatched, expectedIdentity, initialHash)
			if err != nil {
				return nil, err
			}
			return &ApplyResult{Locale: EnglishLocale, LocaleName: "English", BackupPath: restored.BackupPath,
				CleanBackup: p.CleanBackupPath, WasRunning: restored.WasRunning, Forced: restored.Forced, Relaunched: restored.Relaunched}, nil
		}
		locale = &Locale{Language: EnglishLocale, Name: "English", HTMLLang: "en"}
		empty := blockSentinels[0][0] + "\n" + blockSentinels[0][1]
		patch = patchFragments{main: empty, preload: empty}
	} else {
		locale, err = s.store.LoadLocale(localeCode)
		if err != nil {
			return nil, err
		}
		patch, err = s.store.buildLocalizationFragments(locale)
		if err != nil {
			return nil, err
		}
	}
	if features.effortSlider {
		if err := s.store.addEffortSlider(&patch, locale); err != nil {
			return nil, err
		}
	}
	if features.sessionMetrics {
		if err := s.store.addSessionMetrics(&patch, locale); err != nil {
			return nil, err
		}
	}
	s.progress("正在校验补丁", locale.Name)
	// Check the generated fragments before closing the app or touching its files.
	for _, f := range []struct{ name, src string }{
		{"dist/preload.js", patch.preload},
		{"dist/menu.js", patch.menu},
		{"dist/main.js", patch.main},
	} {
		if err := checkJavaScriptSyntax(f.name, f.src); err != nil {
			return nil, err
		}
	}
	wasRunning, err := s.closeApp(p, opts)
	if err != nil {
		return nil, err
	}
	if err := assertArchiveUnchanged(p.AsarPath, initialHash); err != nil {
		return nil, err
	}

	txn := &fileTxn{}
	if err := s.ensureCleanBackup(p, currentPatched, expectedIdentity, initialHash, txn); err != nil {
		return nil, txn.failWithRollback("preparing the clean backup", err)
	}
	s.progress("正在准备干净备份", p.CleanBackupPath)

	backupPath := createUniqueBackupPath(p.ResourcesDir, fmt.Sprintf("app.asar.bak-%s", formatTimestamp()))
	if err := copyFileExclusive(p.AsarPath, backupPath); err != nil {
		return nil, txn.failWithRollback("creating a safety backup", err)
	}
	noteBackupAccumulation(s, p.ResourcesDir)
	s.progress("已创建安全备份", backupPath)

	source := p.AsarPath
	if english {
		// Start with the verified original to remove previous locale changes,
		// including the native menus and secondary windows.
		source = p.CleanBackupPath
	}
	staged, err := buildPatchedArchive(source, locale, patch)
	if err != nil {
		return nil, txn.failWithRollback("building the patched archive", err)
	}
	defer os.Remove(staged)

	// Verify the staged archive in full before it replaces the live one.
	arc, err := ReadArchive(staged)
	if err != nil {
		return nil, txn.failWithRollback("verifying the patched archive", err)
	}
	if patched := detectPatchState(staged); patched == nil || !*patched {
		return nil, txn.failWithRollback("verifying the patched archive",
			fmt.Errorf("patch markers were not found in the archive"))
	}
	if !SameRelease(expectedIdentity, arc.Identity()) {
		return nil, txn.failWithRollback("verifying the patched archive",
			fmt.Errorf("application identity changed in the patched archive"))
	}
	s.progress("正在校验补丁内容", "")
	if code, _, installed, ok := readArchiveLocale(staged); !ok || code != locale.Language || installed != features {
		return nil, txn.failWithRollback("verifying the patched archive",
			fmt.Errorf("the patched archive does not declare the requested locale and features"))
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
	s.progress("正在写回 app.asar", p.AsarPath)

	if err := atomicReplaceFile(staged, p.AsarPath, func(candidate string) error {
		running, err := s.appRunning(p)
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
		Forced:      opts.Force && wasRunning,
	}

	if opts.Restart {
		result.Relaunched = s.relaunch(p)
	}
	return result, nil
}

// readArchiveIdentity verifies the Antigravity identity in package.json.
func readArchiveIdentity(asarPath string) (*Identity, error) {
	arc, err := ReadArchive(asarPath)
	if err != nil {
		return nil, fmt.Errorf("could not read Antigravity archive %s: %w", asarPath, err)
	}
	id := arc.Identity()
	if id == nil || id.Name != "antigravity" {
		return nil, fmt.Errorf("could not verify Antigravity archive identity: %s", asarPath)
	}
	return id, nil
}

type patchFragments struct {
	main          string
	menu          string
	preload       string
	windowPreload string
}

func (s *LocaleStore) buildLocalizationFragments(locale *Locale) (patchFragments, error) {
	preload, err := s.buildPreloadFragment(locale)
	if err != nil {
		return patchFragments{}, err
	}
	menu, err := s.buildMenuFragment(locale)
	if err != nil {
		return patchFragments{}, err
	}
	main, err := s.buildDialogFragment(locale)
	if err != nil {
		return patchFragments{}, err
	}
	return patchFragments{main: main, menu: menu, preload: preload, windowPreload: preload}, nil
}

// buildPatchedArchive produces a patched copy of the archive in a temporary
// directory and returns its path.
func buildPatchedArchive(srcPath string, locale *Locale, patch patchFragments) (string, error) {
	arc, err := ReadArchive(srcPath)
	if err != nil {
		return "", err
	}

	main, err := arc.ReadFile("dist/main.js")
	if err != nil {
		return "", err
	}
	preload, err := arc.ReadFile("dist/preload.js")
	if err != nil {
		return "", err
	}

	var newMain string
	if locale.Language == EnglishLocale {
		newMain = injectFragment(string(main), patch.main, nil, "")
	} else {
		newMain, err = patchMain(string(main), locale, patch.main)
		if err != nil {
			return "", err
		}
	}
	newPreload := patchPreload(string(preload), patch.preload)
	patches := []Patch{
		{Path: "dist/main.js", Data: []byte(newMain)},
		{Path: "dist/preload.js", Data: []byte(newPreload)},
	}
	if locale.Language != EnglishLocale {
		extra, err := buildLocalizedWindowPatches(arc, patch)
		if err != nil {
			return "", err
		}
		patches = append(patches, extra...)
	}

	tmp, err := os.CreateTemp("", "antigravity-i18n-patched-*.asar")
	if err != nil {
		return "", err
	}
	staged := tmp.Name()
	if err := tmp.Close(); err != nil {
		os.Remove(staged)
		return "", err
	}

	err = WritePatched(srcPath, staged, patches)
	if err != nil {
		os.Remove(staged)
		return "", err
	}
	return staged, nil
}

func buildLocalizedWindowPatches(arc *Archive, patch patchFragments) ([]Patch, error) {
	menu, err := arc.ReadFile("dist/menu.js")
	if err != nil {
		return nil, err
	}
	newMenu, err := patchMenu(string(menu), patch.menu)
	if err != nil {
		return nil, err
	}
	patches := []Patch{{Path: "dist/menu.js", Data: []byte(newMenu)}}
	// The IDE installer and SSH password prompt use separate windows and do
	// not load the Hub preload. Older releases can omit either window.
	for _, rel := range []string{"dist/ideInstall/wizardPreload.js", "dist/remote/sshPromptPreload.js"} {
		if arc.lookup(rel) == nil {
			continue
		}
		source, err := arc.ReadFile(rel)
		if err != nil {
			return nil, err
		}
		if rel == "dist/remote/sshPromptPreload.js" {
			const promptPath = "dist/remote/sshPrompts.js"
			prompt, err := arc.ReadFile(promptPath)
			if err != nil {
				return nil, err
			}
			// The message comes from the remote SSH process, so preserve it
			// even when it happens to match a translated button label.
			const original = `<div id="message"></div>`
			const protected = `<div id="message" translate="no"></div>`
			if !strings.Contains(string(prompt), original) && !strings.Contains(string(prompt), protected) {
				return nil, fmt.Errorf("could not locate the SSH prompt message container in %s", promptPath)
			}
			promptSource := strings.ReplaceAll(string(prompt), original, protected)
			if err := checkJavaScriptSyntax(promptPath, promptSource); err != nil {
				return nil, err
			}
			patches = append(patches, Patch{Path: promptPath, Data: []byte(promptSource)})
		}
		patched := patchPreload(string(source), patch.windowPreload)
		if err := checkJavaScriptSyntax(rel, patched); err != nil {
			return nil, err
		}
		patches = append(patches, Patch{Path: rel, Data: []byte(patched)})
	}

	return patches, nil
}

// RestoreOfficial restores the official untranslated archive.
func (s *Service) RestoreOfficial(opts Options) (*RestoreResult, error) {
	s.clearNotes()

	p, err := resolveAppPaths(opts.AppDir)
	if err != nil {
		return nil, err
	}
	unlock, err := lockInstallation(p.ResourcesDir)
	if err != nil {
		return nil, err
	}
	defer unlock()

	initialHash, err := hashArchive(p.AsarPath)
	if err != nil {
		return nil, err
	}
	expectedIdentity, err := readArchiveIdentity(p.AsarPath)
	if err != nil {
		return nil, err
	}
	currentPatched, err := detectPatchInArchive(p.AsarPath)
	if err != nil {
		return nil, err
	}

	return s.restoreOfficial(p, opts, currentPatched, expectedIdentity, initialHash)
}

func (s *Service) restoreOfficial(p *paths, opts Options, currentPatched bool, expectedIdentity *Identity, initialHash string) (*RestoreResult, error) {
	if currentPatched {
		s.progress("正在定位干净备份", p.CleanBackupPath)
	}
	txn := &fileTxn{}
	if err := s.ensureCleanBackup(p, currentPatched, expectedIdentity, initialHash, txn); err != nil {
		return nil, txn.failWithRollback("locating a clean English archive", err)
	}
	if !currentPatched {
		if err := assertArchiveUnchanged(p.AsarPath, initialHash); err != nil {
			return nil, txn.failWithRollback("refreshing the clean backup", err)
		}
		writeState(p, EnglishLocale)
		s.progress("已经是官方英文版本", "")
		return &RestoreResult{AlreadyEnglish: true}, nil
	}

	backupHash, err := hashArchive(p.CleanBackupPath)
	if err != nil {
		return nil, txn.failWithRollback("reading the clean backup", err)
	}
	wasRunning, err := s.closeApp(p, opts)
	if err != nil {
		return nil, txn.failWithRollback("closing Antigravity", err)
	}
	if err := assertArchiveUnchanged(p.AsarPath, initialHash); err != nil {
		return nil, txn.failWithRollback("checking the installed archive", err)
	}

	backupPath := createUniqueBackupPath(p.ResourcesDir,
		fmt.Sprintf("app.asar.bak-before-restore-%s", formatTimestamp()))
	if err := copyFileExclusive(p.AsarPath, backupPath); err != nil {
		return nil, txn.failWithRollback("creating a safety backup", err)
	}
	noteBackupAccumulation(s, p.ResourcesDir)
	s.progress("已创建安全备份", backupPath)

	s.progress("正在写回 app.asar", p.AsarPath)
	if err := atomicReplaceFile(p.CleanBackupPath, p.AsarPath, func(candidate string) error {
		running, err := s.appRunning(p)
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

	writeState(p, EnglishLocale)

	result := &RestoreResult{
		BackupPath: backupPath,
		WasRunning: wasRunning,
		Forced:     opts.Force && wasRunning,
	}
	if opts.Restart {
		result.Relaunched = s.relaunch(p)
	}
	return result, nil
}

// A launch error does not undo a completed patch, but must not be reported as
// a successful restart either.
func (s *Service) relaunch(p *paths) bool {
	if err := s.launchApp(p); err != nil {
		s.notes = append(s.notes, fmt.Sprintf("Could not relaunch Antigravity: %v. Open it manually.", err))
		return false
	}
	return true
}

// An unpatched installation supplies the exact archive to restore, even if a
// different clean build declares the same version.
func (s *Service) ensureCleanBackup(p *paths, patched bool, identity *Identity, initialHash string, txn *fileTxn) error {
	state := inspectCleanBackup(p.CleanBackupPath, identity)
	source := p.AsarPath
	if patched {
		if state == cleanBackupValid {
			return nil
		}
		var err error
		source, err = findPristineArchive(p.ResourcesDir, identity)
		if err != nil {
			return err
		}
		if source == "" {
			return fmt.Errorf("no verified pristine English app.asar was found. Reinstall or update Antigravity, then try again")
		}
	} else if state == cleanBackupValid {
		if hash, err := hashArchive(p.CleanBackupPath); err == nil && hash == initialHash {
			return nil
		}
	}

	if err := txn.replace(source, p.CleanBackupPath, func(staged string) error {
		if !patched {
			if err := assertArchiveUnchanged(staged, initialHash); err != nil {
				return err
			}
		}
		if status := inspectCleanBackup(staged, identity); status != cleanBackupValid {
			return fmt.Errorf("the candidate clean archive is %s: %s", status, source)
		}
		return nil
	}); err != nil {
		return err
	}
	if patched {
		s.notes = append(s.notes, fmt.Sprintf("Recovered a pristine archive from %s", filepath.Base(source)))
	} else if state != cleanBackupMissing {
		s.notes = append(s.notes, "Refreshed the clean backup from the current unpatched archive.")
	}
	return nil
}
