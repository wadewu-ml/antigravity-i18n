//go:build packedlocales

package main

import (
	"archive/zip"
	_ "embed"
	"io/fs"
	"strings"
)

//go:embed build/generated/locales.zip
var packedLocales string

func openLocaleAssets() (fs.FS, error) {
	return zip.NewReader(strings.NewReader(packedLocales), int64(len(packedLocales)))
}
