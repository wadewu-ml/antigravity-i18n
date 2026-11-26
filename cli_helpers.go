package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"regravity/internal/core"
)

// displayLocales lists the locale packs with coverage, for the CLI.
func displayLocales(svc *core.Service) []core.LocaleInfo {
	locales, err := svc.DescribeLocales()
	if err != nil {
		return nil
	}
	return locales
}

func fail(msg string) int {
	fmt.Fprintln(os.Stderr, "error:", msg)
	return 1
}

// confirmShutdown asks before the tool closes a running application. Patching
// app.asar requires the app to be closed, and doing that without asking would
// discard unsaved work.
func confirmShutdown(what string, force bool) bool {
	fmt.Println(what)
	if force {
		fmt.Println("If Antigravity is running, it will be force-closed. Unsaved work may be lost.")
	} else {
		fmt.Println("If Antigravity is running, it will be asked to close. The operation stops if it stays open.")
	}
	fmt.Print("Continue? [y/N] ")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil && line == "" {
		return false
	}
	answer := strings.ToLower(strings.TrimSpace(line))
	return answer == "y" || answer == "yes"
}
