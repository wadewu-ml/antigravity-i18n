package core

import (
	"os"
	"testing"
)

// TestAllLocalesValidate loads and validates every shipped locale through the
// same path the installer uses, so a malformed pack fails the build rather than
// a user's apply.
func TestAllLocalesValidate(t *testing.T) {
	assets := os.DirFS("../..")
	store, err := NewLocaleStore(assets, assets, "assets/locales", "assets/patches")
	if err != nil {
		t.Fatal(err)
	}
	list := store.ListLocales()
	if len(list) != 9 {
		t.Fatalf("expected 9 locales, got %d: %v", len(list), list)
	}
	for _, code := range list {
		loc, err := store.LoadLocale(code)
		if err != nil {
			t.Fatalf("locale %s: %v", code, err)
		}
		if len(loc.Text) == 0 {
			t.Fatalf("locale %s has no text", code)
		}
		t.Logf("OK %-8s name=%-14s entries=%d patterns=%d menu=%d dialogs=%d",
			code, loc.Name, len(loc.Text), len(loc.Patterns), len(loc.Menu), len(loc.Dialogs))
	}
}

// Fragments must come out as valid JavaScript: they are injected into the app's
// own scripts, and a syntax error there stops Antigravity from starting.
func TestFragmentsBuild(t *testing.T) {
	assets := os.DirFS("../..")
	store, err := NewLocaleStore(assets, assets, "assets/locales", "assets/patches")
	if err != nil {
		t.Fatal(err)
	}
	for _, code := range store.ListLocales() {
		loc, err := store.LoadLocale(code)
		if err != nil {
			t.Fatal(err)
		}
		engine, err := store.buildPreloadFragment(loc)
		if err != nil {
			t.Fatalf("%s: %v", code, err)
		}
		if len(engine) < 1000 {
			t.Fatalf("%s: engine fragment suspiciously small: %d", code, len(engine))
		}
		if err := checkJavaScriptSyntax("dist/preload.js", engine); err != nil {
			t.Fatalf("%s engine: %v", code, err)
		}
		menu, err := store.buildMenuFragment(loc)
		if err != nil {
			t.Fatalf("%s: %v", code, err)
		}
		if err := checkJavaScriptSyntax("dist/menu.js", menu); err != nil {
			t.Fatalf("%s menu: %v", code, err)
		}
		dialog, err := store.buildDialogFragment(loc)
		if err != nil {
			t.Fatalf("%s: %v", code, err)
		}
		if err := checkJavaScriptSyntax("dist/main.js", dialog); err != nil {
			t.Fatalf("%s dialog: %v", code, err)
		}
		t.Logf("OK %-8s engine=%d menu=%d dialog=%d", code, len(engine), len(menu), len(dialog))
	}
}
