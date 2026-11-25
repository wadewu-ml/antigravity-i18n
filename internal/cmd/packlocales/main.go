package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"antigravity-i18n/internal/core"
)

func main() {
	if err := packLocales(); err != nil {
		fmt.Fprintln(os.Stderr, "cannot package the locales:", err)
		os.Exit(1)
	}
}

func packLocales() error {
	assets := os.DirFS(".")
	store, err := core.NewLocaleStore(assets, assets, "assets/locales", "assets/patches")
	if err != nil {
		return err
	}
	dir := filepath.Join("build", "generated")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	file, err := os.CreateTemp(dir, ".locales-*.zip")
	if err != nil {
		return err
	}
	defer os.Remove(file.Name())
	writeErr := store.WriteLocaleBundle(file)
	if err := errors.Join(writeErr, file.Close()); err != nil {
		return err
	}
	return os.Rename(file.Name(), filepath.Join(dir, "locales.zip"))
}
