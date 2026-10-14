package main

import (
	"fmt"
	"strconv"
	"strings"

	"antigravity-i18n/internal/core"
)

const usage = `antigravity-i18n - localize the Antigravity desktop interface

Usage:
  antigravity-i18n [command] [options]

Commands:
  apply      Install a language pack into Antigravity
  restore    Restore the official English version
  status     Report the current localization state
  locales    List the installed language packs
  gui        Open the graphical interface (default when no command is given)

Options:
  --app-dir <path>   Antigravity installation directory
  --locale <code>    Language pack to install (default: zh-CN)
  --no-restart       Do not relaunch Antigravity afterwards
  --force            Skip the graceful shutdown and force Antigravity to close
  --timeout <ms>     How long to wait for a graceful shutdown (default: 30000)
  --yes              Assume yes for the shutdown confirmation
  -h, --help         Show this help
  -v, --version      Show the tool version
`

// runCLI returns the process exit code.
func runCLI(svc *core.Service, args []string) int {
	cmd := args[0]
	rest := args[1:]

	opts := core.Options{
		Restart: true,
	}
	var confirm bool

	for i := 0; i < len(rest); i++ {
		arg := rest[i]
		switch arg {
		case "--app-dir":
			if i+1 >= len(rest) {
				return fail("--app-dir needs a path")
			}
			i++
			opts.AppDir = rest[i]
		case "--locale":
			if i+1 >= len(rest) {
				return fail("--locale needs a language code")
			}
			i++
			opts.Locale = rest[i]
		case "--no-restart":
			opts.Restart = false
		case "--force":
			opts.Force = true
		case "--yes", "-y":
			confirm = true
		case "--timeout":
			if i+1 >= len(rest) {
				return fail("--timeout needs a number of milliseconds")
			}
			i++
			ms, err := strconv.Atoi(rest[i])
			if err != nil {
				return fail("--timeout needs a number of milliseconds")
			}
			opts.TimeoutMs = ms
		case "-h", "--help":
			fmt.Print(usage)
			return 0
		case "-v", "--version":
			fmt.Println("antigravity-i18n " + core.Version)
			return 0
		default:
			if strings.HasPrefix(arg, "--app-dir=") {
				opts.AppDir = strings.TrimPrefix(arg, "--app-dir=")
				continue
			}
			if strings.HasPrefix(arg, "--locale=") {
				opts.Locale = strings.TrimPrefix(arg, "--locale=")
				continue
			}
			if strings.HasPrefix(arg, "--timeout=") {
				ms, err := strconv.Atoi(strings.TrimPrefix(arg, "--timeout="))
				if err != nil {
					return fail("--timeout needs a number of milliseconds")
				}
				opts.TimeoutMs = ms
				continue
			}
			return fail(fmt.Sprintf("unknown option %q\n\n%s", arg, usage))
		}
	}

	opts.Progress = func(stage, detail string) {
		if detail != "" {
			fmt.Printf("  %s: %s\n", stage, detail)
			return
		}
		fmt.Printf("  %s\n", stage)
	}

	switch cmd {
	case "apply":
		return cliApply(svc, opts, confirm)
	case "restore":
		return cliRestore(svc, opts, confirm)
	case "status":
		return cliStatus(svc, opts)
	case "locales":
		return cliLocales(svc)
	case "gui":
		runWails(svc)
		return 0
	default:
		return fail(fmt.Sprintf("unknown command %q\n\n%s", cmd, usage))
	}
}

func cliApply(svc *core.Service, opts core.Options, confirmed bool) int {
	if !confirmed {
		if !confirmShutdown("Applying a language pack rewrites app.asar and Antigravity must not be running.") {
			return 1
		}
	}
	result, err := svc.ApplyLocale(opts)
	for _, note := range svc.Notes() {
		fmt.Println("  note:", note)
	}
	if err != nil {
		return fail(err.Error())
	}
	fmt.Printf("Successfully switched to %s (%s)\n", result.LocaleName, result.Locale)
	fmt.Printf("  Backup saved at: %s\n", result.BackupPath)
	if opts.Restart {
		fmt.Println("  Antigravity has been relaunched.")
	}
	return 0
}

func cliRestore(svc *core.Service, opts core.Options, confirmed bool) int {
	if !confirmed {
		if !confirmShutdown("Restoring the English version rewrites app.asar and Antigravity must not be running.") {
			return 1
		}
	}
	result, err := svc.RestoreOfficial(opts)
	for _, note := range svc.Notes() {
		fmt.Println("  note:", note)
	}
	if err != nil {
		return fail(err.Error())
	}
	if result.AlreadyEnglish {
		fmt.Println("Antigravity is already using the official English archive.")
		return 0
	}
	fmt.Println("Successfully switched back to the official English version.")
	if result.BackupPath != "" {
		fmt.Printf("  A copy of the localized build was saved at: %s\n", result.BackupPath)
	}
	if opts.Restart {
		fmt.Println("  Antigravity has been relaunched.")
	}
	return 0
}

func cliStatus(svc *core.Service, opts core.Options) int {
	status, err := svc.GetStatus(opts.AppDir)
	if err != nil {
		return fail(err.Error())
	}
	fmt.Printf("App directory : %s\n", status.AppDir)
	fmt.Printf("Archive       : %s\n", status.AsarPath)
	fmt.Printf("Language      : %s", status.CurrentLanguage)
	if status.LocaleName != "" {
		fmt.Printf(" (%s)", status.LocaleName)
	}
	fmt.Println()
	fmt.Printf("Clean backup  : %s (%s)\n", status.CleanBackupStatus, status.CleanBackupPath)
	if status.LastPatchedAt != nil {
		fmt.Printf("Last patched  : %s\n", *status.LastPatchedAt)
	}
	if status.StateIsStale {
		fmt.Println("Note          : the recorded state marker is stale; the archive itself was read directly.")
	}
	for _, locale := range displayLocales(svc) {
		marker := "  "
		if locale.Code == status.CurrentLanguage {
			marker = "->"
		}
		fmt.Printf("  %s %-8s %-24s %d translations, %d%% coverage\n",
			marker, locale.Code, locale.Name, locale.Translations, locale.Coverage)
	}
	return 0
}

func cliLocales(svc *core.Service) int {
	for _, locale := range displayLocales(svc) {
		fmt.Printf("%-8s %-24s %-24s %5d translations  %3d%% coverage\n",
			locale.Code, locale.Name, locale.EnglishName, locale.Translations, locale.Coverage)
	}
	return 0
}
