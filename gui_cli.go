//go:build cli

package main

import (
	"fmt"

	"regravity/internal/core"
)

const guiHelp = ""

func runWails(*core.Service) { fmt.Print(usage) }
