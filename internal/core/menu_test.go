package core

import (
	"os"
	"testing"

	"github.com/dop251/goja"
)

// The WSL menu is populated after the initial menu has already been localized.
// Exercise the resulting JavaScript so a translated File label cannot prevent
// insertion and a later setApplicationMenu cannot expose untranslated items.
func TestPatchedMenuTranslatesAsynchronousWslItems(t *testing.T) {
	assets := os.DirFS("../..")
	store, err := NewLocaleStore(assets, assets, "assets/locales", "assets/patches")
	if err != nil {
		t.Fatal(err)
	}
	loc, err := store.LoadLocale(DefaultLocale)
	if err != nil {
		t.Fatal(err)
	}
	fragment, err := store.buildMenuFragment(loc)
	if err != nil {
		t.Fatal(err)
	}
	source := `
const installs = [];
const menu = {items: [{label: 'File', submenu: {items: [], insert(index, item) { this.items.splice(index, 0, item); }}}]};
const electron_1 = {Menu: {setApplicationMenu(value) { installs.push(JSON.parse(JSON.stringify(value))); }}};
function setupApplicationMenu() {
    // Re-apply the menu so the change takes effect.
    electron_1.Menu.setApplicationMenu(menu);
    void addWslConnectMenu(menu);
}
async function addWslConnectMenu(menu) {
    await Promise.resolve();
    addItemToSubmenu(menu, 'File', 0, {label: 'Connect to WSL', submenu: {items: [{label: 'Ubuntu'}]}});
    addItemToSubmenu(menu, 'File', 1, {label: 'Reopen Locally'});
    electron_1.Menu.setApplicationMenu(menu);
}
function addItemToSubmenu(appMenu, submenuLabel, position, item) {
    const submenuItem = appMenu.items.find((item) => item.label === submenuLabel);
    if (!submenuItem?.submenu) return;
    submenuItem.submenu.insert(position, item);
}
`
	patched, err := patchMenu(source, fragment)
	if err != nil {
		t.Fatal(err)
	}
	again, err := patchMenu(patched, fragment)
	if err != nil || again != patched {
		t.Fatalf("menu patch is not idempotent: %v", err)
	}
	runtime := goja.New()
	if _, err := runtime.RunString(patched + "\nsetupApplicationMenu();"); err != nil {
		t.Fatal(err)
	}
	checks := map[string]string{
		"String(installs.length)":                                      "2",
		"installs[0].items[0].label":                                   loc.Menu["File"],
		"installs[1].items[0].submenu.items[0].label":                  loc.Menu["Connect to WSL"],
		"installs[1].items[0].submenu.items[1].label":                  loc.Menu["Reopen Locally"],
		"installs[1].items[0].submenu.items[0].submenu.items[0].label": "Ubuntu",
	}
	for expr, want := range checks {
		got, err := runtime.RunString(expr)
		if err != nil {
			t.Fatalf("%s: %v", expr, err)
		}
		if got.String() != want {
			t.Errorf("%s = %q, want %q", expr, got.String(), want)
		}
	}
}
