package core

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// stateMarkerName is the file the tool writes to record which language it last
// installed.
const stateMarkerName = "regravity-state.json"

// legacyStateMarkerNames are markers written before the tool was renamed. They
// are read as a fallback so an existing install does not lose its recorded
// language, and removed once the current marker is in place.
var legacyStateMarkerNames = []string{"antigravity-i18n-state.json", "polygravity-state.json", "antigravity-zh-state.json"}

const (
	antigravityExe         = "Antigravity.exe"
	gracefulTimeoutMs      = 30_000
	gracefulPollIntervalMs = 500
)

// paths addresses the installation the tool operates on.
type paths struct {
	AppDir          string
	ResourcesDir    string
	AsarPath        string
	CleanBackupPath string
	StatePath       string
}

// PossibleAppDirs lists where Antigravity may be installed, in priority order.
// ANTIGRAVITY_APP_DIR wins so a portable or non-default install can be targeted.
func PossibleAppDirs() []string {
	var dirs []string
	if v := os.Getenv("ANTIGRAVITY_APP_DIR"); v != "" {
		dirs = append(dirs, v)
	}
	localAppData := os.Getenv("LOCALAPPDATA")
	if localAppData == "" {
		if home, err := os.UserHomeDir(); err == nil {
			localAppData = filepath.Join(home, "AppData", "Local")
		}
	}
	programFiles := os.Getenv("ProgramFiles")
	if programFiles == "" {
		programFiles = `C:\Program Files`
	}
	programFilesX86 := os.Getenv("ProgramFiles(x86)")
	if programFilesX86 == "" {
		programFilesX86 = `C:\Program Files (x86)`
	}
	if localAppData != "" {
		dirs = append(dirs, filepath.Join(localAppData, "Programs", "antigravity"))
	}
	dirs = append(dirs,
		filepath.Join(programFiles, "Antigravity"),
		filepath.Join(programFilesX86, "Antigravity"))
	return dirs
}

// resolveAppPaths finds the installation, or reports why none was found.
func resolveAppPaths(customAppDir string) (*paths, error) {
	var appDir string
	if customAppDir != "" {
		appDir = customAppDir
	} else {
		for _, candidate := range PossibleAppDirs() {
			if dirExists(candidate) && fileExists(filepath.Join(candidate, "resources", "app.asar")) {
				appDir = candidate
				break
			}
			if dirExists(candidate) && fileExists(filepath.Join(candidate, "app.asar")) {
				appDir = candidate
				break
			}
		}
	}
	if appDir == "" || !dirExists(appDir) {
		return nil, fmt.Errorf(
			"could not locate Antigravity installation. specify the path with --app-dir <path> " +
				"or set ANTIGRAVITY_APP_DIR")
	}

	asarPath := filepath.Join(appDir, "resources", "app.asar")
	if !fileExists(asarPath) {
		alt := filepath.Join(appDir, "app.asar")
		if fileExists(alt) {
			asarPath = alt
			parent := filepath.Dir(filepath.Clean(appDir))
			if strings.EqualFold(filepath.Base(appDir), "resources") && fileExists(filepath.Join(parent, antigravityExe)) {
				appDir = parent
			}
		}
	}
	if !fileExists(asarPath) {
		return nil, fmt.Errorf("Antigravity app.asar was not found or is not a file: %s", asarPath)
	}

	resourcesDir := filepath.Dir(asarPath)
	return &paths{
		AppDir:          appDir,
		ResourcesDir:    resourcesDir,
		AsarPath:        asarPath,
		CleanBackupPath: filepath.Join(resourcesDir, "app.asar.clean-backup"),
		StatePath:       filepath.Join(resourcesDir, stateMarkerName),
	}, nil
}

func dirExists(p string) bool {
	info, err := os.Stat(p)
	return err == nil && info.IsDir()
}

func fileExists(p string) bool {
	info, err := os.Stat(p)
	return err == nil && info.Mode().IsRegular()
}

// statePayload is what the marker records. asarSize is a staleness hint, not the
// source of truth: the archive itself always wins.
type statePayload struct {
	Language    string `json:"language"`
	PatchedAt   string `json:"patchedAt"`
	AsarSize    *int64 `json:"asarSize"`
	ToolVersion string `json:"toolVersion"`
}

// readState reads the current marker, falling back to a pre-rename one. A
// missing or unreadable marker means "unknown", never "English".
func readState(p *paths) *statePayload {
	candidates := append([]string{p.StatePath}, legacyPaths(p)...)
	for _, candidate := range candidates {
		data, err := os.ReadFile(candidate)
		if err != nil {
			continue
		}
		var parsed statePayload
		if err := json.Unmarshal(data, &parsed); err != nil {
			continue
		}
		if parsed.Language != "" {
			return &parsed
		}
	}
	return nil
}

func legacyPaths(p *paths) []string {
	out := make([]string, 0, len(legacyStateMarkerNames))
	for _, name := range legacyStateMarkerNames {
		out = append(out, filepath.Join(filepath.Dir(p.StatePath), name))
	}
	return out
}

// writeState records what language was just applied.
func writeState(p *paths, language string) {
	payload := statePayload{
		Language:    language,
		PatchedAt:   time.Now().UTC().Format(time.RFC3339),
		ToolVersion: Version,
	}
	if info, err := os.Stat(p.AsarPath); err == nil {
		size := info.Size()
		payload.AsarSize = &size
	}

	data, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return
	}
	data = append(data, '\n')
	if err := os.WriteFile(p.StatePath, data, 0o644); err != nil {
		// A missing marker degrades detection but must never fail the patch.
		return
	}
	// A stale pre-rename marker beside the new one would keep being read as a
	// fallback, so it is removed once the new marker is safely written.
	for _, legacy := range legacyPaths(p) {
		if legacy != p.StatePath {
			_ = os.Remove(legacy)
		}
	}
}

// Version is the tool version reported by the CLI and the window. Release
// builds stamp the real number in with -ldflags (see build.ps1); an unstamped
// local build keeps "dev" rather than claiming a release number it is not.
var Version = "dev"

// launchAntigravity starts the app again after a patch.
func launchAntigravity(p *paths) error {
	exe := filepath.Join(p.AppDir, antigravityExe)
	if !fileExists(exe) {
		return fmt.Errorf("Antigravity.exe was not found in %s", p.AppDir)
	}
	cmd := exec.Command(exe)
	if err := cmd.Start(); err != nil {
		return err
	}
	if cmd.Process != nil {
		_ = cmd.Process.Release()
	}
	return nil
}
