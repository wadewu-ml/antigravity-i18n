package core

import (
	"archive/zip"
	"bytes"
	"io/fs"
	"os"
	"reflect"
	"testing"
)

func TestPackedLocalesMatchSources(t *testing.T) {
	source := newTestService(t).store
	var output bytes.Buffer
	if err := source.WriteLocaleBundle(&output); err != nil {
		t.Fatal(err)
	}
	archive, err := zip.NewReader(bytes.NewReader(output.Bytes()), int64(output.Len()))
	if err != nil {
		t.Fatal(err)
	}
	assets := os.DirFS("../..")
	packed, err := NewLocaleStore(archive, assets, source.localeDir, source.patchDir)
	if err != nil {
		t.Fatal(err)
	}
	want, err := source.Describe()
	if err != nil {
		t.Fatal(err)
	}
	got, err := packed.Describe()
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("packed metadata differs from the sources: %+v, %v", got, err)
	}
	for _, locale := range want {
		name := source.localeDir + "/" + locale.Code + ".json"
		raw, err := fs.ReadFile(assets, name)
		if err != nil {
			t.Fatal(err)
		}
		unpacked, err := fs.ReadFile(archive, name)
		if err != nil || !bytes.Equal(raw, unpacked) {
			t.Fatalf("packed dictionary differs from %s: %v", name, err)
		}
		if _, err := packed.LoadLocale(locale.Code); err != nil {
			t.Fatal(err)
		}
	}
}

func TestBundledLocales(t *testing.T) {
	store := newTestService(t).store
	for _, code := range store.ListLocales() {
		t.Run(code, func(t *testing.T) {
			loc, err := store.LoadLocale(code)
			if err != nil {
				t.Fatal(err)
			}
			if (code == "zh-CN" || code == "zh-Hant") && loc.Text["Turbo"] != "Turbo" {
				t.Fatal("the permission preset name Turbo must be preserved")
			}
			for _, fragment := range []struct {
				path  string
				build func(*Locale) (string, error)
			}{
				{"dist/preload.js", store.buildPreloadFragment},
				{"dist/menu.js", store.buildMenuFragment},
				{"dist/main.js", store.buildDialogFragment},
			} {
				source, err := fragment.build(loc)
				if err != nil {
					t.Fatal(err)
				}
				if err := checkJavaScriptSyntax(fragment.path, source); err != nil {
					t.Fatal(err)
				}
			}
		})
	}
}
