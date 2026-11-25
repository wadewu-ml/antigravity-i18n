//go:build !packedlocales

package main

import (
	"embed"
	"io/fs"
)

// Ordinary Go builds and tests work directly from a fresh checkout.
//
//go:embed assets/locales/*.json
var sourceLocales embed.FS

func openLocaleAssets() (fs.FS, error) { return sourceLocales, nil }
