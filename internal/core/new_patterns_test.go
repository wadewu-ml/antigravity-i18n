package core

import (
	"os"
	"strings"
	"testing"

	"github.com/dop251/goja"
)

// The renderer interpolates account names, shortcuts and downloaded filenames
// into accessibility labels. The wrappers should translate without interpreting
// dollar replacement tokens or changing the interpolated values.
func TestEngineAntigravity232DynamicLabels(t *testing.T) {
	assets := os.DirFS("../..")
	store, err := NewLocaleStore(assets, assets, "assets/locales", "assets/patches")
	if err != nil {
		t.Fatal(err)
	}
	expected := map[string]struct {
		auth, download string
	}{
		"zh-CN":   {"Example $& Account 的授权码", "在文件夹中显示 report $&.pdf"},
		"zh-Hant": {"Example $& Account 的授權碼", "在資料夾中顯示 report $&.pdf"},
		"ja":      {"Example $& Account の認証コード", "report $&.pdf をフォルダーに表示"},
		"ko":      {"Example $& Account의 인증 코드", "폴더에서 report $&.pdf 표시"},
		"de":      {"Autorisierungscode für Example $& Account", "report $&.pdf im Ordner anzeigen"},
		"es":      {"Código de autorización de Example $& Account", "Mostrar report $&.pdf en la carpeta"},
		"fr":      {"Code d’autorisation pour Example $& Account", "Afficher report $&.pdf dans le dossier"},
		"pt-BR":   {"Código de autorização de Example $& Account", "Mostrar report $&.pdf na pasta"},
		"ru":      {"Код авторизации для Example $& Account", "Показать report $&.pdf в папке"},
	}
	for _, code := range store.ListLocales() {
		t.Run(code, func(t *testing.T) {
			loc, err := store.LoadLocale(code)
			if err != nil {
				t.Fatal(err)
			}
			engine, err := store.buildPreloadFragment(loc)
			if err != nil {
				t.Fatal(err)
			}
			runtime := goja.New()
			if _, err := runtime.RunString(`const document = { readyState: 'loading', addEventListener() {} };`); err != nil {
				t.Fatal(err)
			}
			if _, err := runtime.RunString(engine); err != nil {
				t.Fatal(err)
			}
			translate, ok := goja.AssertFunction(runtime.Get("translateString"))
			if !ok {
				t.Fatal("engine does not expose translateString")
			}
			translated := func(input string) string {
				t.Helper()
				result, err := translate(goja.Undefined(), runtime.ToValue(input))
				if err != nil {
					t.Fatal(err)
				}
				return result.String()
			}
			for _, input := range []string{"Authorization code for Example $& Account", "Auth code for Example $& Account"} {
				if got := translated(input); got != expected[code].auth {
					t.Errorf("translateString(%q) = %q, want %q", input, got, expected[code].auth)
				}
			}
			if got := translated("Show report $&.pdf in folder"); got != expected[code].download {
				t.Errorf("filename label = %q, want %q", got, expected[code].download)
			}
			for _, tc := range []struct {
				input string
				keep  []string
			}{
				{"Recording shortcut for Custom $& Action. Press keys to bind, Enter to save, Escape to cancel.", []string{"Custom $& Action"}},
				{"Shortcut Ctrl+Shift+$& for Custom $& Action", []string{"Ctrl+Shift+$&", "Custom $& Action"}},
				{"Remove shortcut Ctrl+Shift+$& for Custom $& Action", []string{"Ctrl+Shift+$&", "Custom $& Action"}},
				{"Shortcut Ctrl+K for Search for plugins", []string{"Ctrl+K", "Search for plugins"}},
				{"Remove shortcut Ctrl+K for Search for plugins", []string{"Ctrl+K", "Search for plugins"}},
				{"Add shortcut for Custom $& Action", []string{"Custom $& Action"}},
				{"Reset Custom $& Action to default", []string{"Custom $& Action"}},
				{"report $&.pdf download progress", []string{"report $&.pdf"}},
				{"Remove report $&.pdf from list", []string{"report $&.pdf"}},
				{"Go to turn 2 of 12", []string{"2", "12"}},
				{"Voice Speed, 1.5x", []string{"1.5x"}},
				{"Voice speed, 1.5x", []string{"1.5x"}},
				{"More actions for Example $& Project", []string{"Example $& Project"}},
				{"Search for \"Example $& Plugin\" in the Marketplace", []string{"Example $& Plugin"}},
				{"No plugins found for \"Example $& Plugin\".", []string{"Example $& Plugin"}},
				{"No MCP servers found for \"Example $& Server\"", []string{"Example $& Server"}},
				{"Switch environment (Example $& Host)", []string{"Example $& Host"}},
				{"Open Subagent: Example $& Agent", []string{"Example $& Agent"}},
				{"Search Google for \"Example $& Query\"", []string{"Example $& Query"}},
				{"user+$&@example.com asks:", []string{"user+$&@example.com"}},
			} {
				got := translated(tc.input)
				if got == tc.input {
					t.Errorf("dynamic label was not translated: %q", tc.input)
				}
				for _, captured := range tc.keep {
					if !strings.Contains(got, captured) {
						t.Errorf("translation %q changed the interpolated value %q", got, captured)
					}
				}
			}
			for _, input := range []string{
				"Default needs the terminal sandbox, which isn't available on this host.",
				"Auto needs the terminal sandbox, which isn't available on this host.",
				"Default and Auto need the terminal sandbox, which isn't available on this host.",
			} {
				got := translated(input)
				if got == input || strings.Contains(got, "the terminal sandbox") {
					t.Errorf("sandbox availability explanation was not translated: %q", got)
				}
			}
			// Unknown modes and incomplete labels are not word-by-word replacements.
			for _, input := range []string{
				"Turbo needs the terminal sandbox, which isn't available on this host.",
				"Default needs the terminal sandbox, which isn't available on this host. More details follow.",
				"Authorization code for",
				"Voice speed, tomorrow",
			} {
				if got := translated(input); got != input {
					t.Errorf("near-match %q unexpectedly translated to %q", input, got)
				}
			}
		})
	}
}
