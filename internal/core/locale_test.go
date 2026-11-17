package core

import "testing"

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
