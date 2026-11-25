package main

import "embed"

// The patch templates and translation engine are embedded alongside the locale
// packs, so the tool has no dependency on Node or the source directory.
//
//go:embed assets/patches
var patchesFS embed.FS

// frontendFS holds the hand-written static window: HTML, CSS and JS only. Wails
// serves it as-is, which is why no Node toolchain is involved at build time.
//
//go:embed all:frontend/dist
var frontendFS embed.FS
