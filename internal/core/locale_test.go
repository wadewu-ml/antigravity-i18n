package core

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"io/fs"
	"maps"
	"os"
	"reflect"
	"slices"
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

// Coverage in the UI counts only static text. Compare the other translation
// surfaces here so a missing native label or dynamic rule cannot ship unnoticed.
func TestLanguagePackParity(t *testing.T) {
	store := newTestService(t).store
	rawLabels, err := fs.ReadFile(store.patchesFS, store.patchDir+"/metrics-labels.json")
	if err != nil {
		t.Fatal(err)
	}
	var metrics map[string]map[string]string
	if err := json.Unmarshal(rawLabels, &metrics); err != nil {
		t.Fatal(err)
	}
	reference, err := store.LoadLocale(DefaultLocale)
	if err != nil {
		t.Fatal(err)
	}
	for _, code := range store.ListLocales() {
		t.Run(code, func(t *testing.T) {
			loc, err := store.LoadLocale(code)
			if err != nil {
				t.Fatal(err)
			}
			checkKeys := func(section string, want, got map[string]string) {
				t.Helper()
				if !slices.Equal(slices.Sorted(maps.Keys(want)), slices.Sorted(maps.Keys(got))) {
					t.Errorf("%s has different source keys", section)
				}
			}
			checkKeys("session metrics", metrics["en"], metrics[code])
			for key, label := range metrics[code] {
				if label == "" {
					t.Errorf("empty session metrics label %s", key)
				}
			}
			checkKeys("text", reference.Text, loc.Text)
			checkKeys("menu", reference.Menu, loc.Menu)
			checkKeys("dialogs", reference.Dialogs, loc.Dialogs)
			for group, want := range reference.ValueMaps {
				checkKeys(group, want, loc.ValueMaps[group])
			}
			allowed := make(map[string]bool, len(loc.AllowSourceEqual))
			for _, key := range loc.AllowSourceEqual {
				allowed[key] = true
			}
			for source, translation := range loc.Text {
				if source == translation && !allowed[source] {
					t.Errorf("unreviewed English text %q", source)
				}
			}
			// Languages may need additional rules for their own grammar.
			rules := map[string]bool{}
			for _, rule := range loc.Patterns {
				rules[rule.ID] = true
			}
			for _, rule := range reference.Patterns {
				if !rules[rule.ID] {
					t.Errorf("missing dynamic rule %s", rule.ID)
				}
			}
			inline := map[string]string{}
			for _, message := range loc.InlineMessages {
				inline[message.ID] = message.Source
			}
			for _, message := range reference.InlineMessages {
				if inline[message.ID] != message.Source {
					t.Errorf("missing inline source %s", message.ID)
				}
			}
		})
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
