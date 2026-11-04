package core

import (
	"strings"
	"testing"
)

// Antigravity 2.29.0 still has no language switch of its own. The remote
// debugging port block is the anchor applyLanguageSwitch inserts beside.
func TestLanguageSwitchAnchorsOnAntigravity229Main(t *testing.T) {
	source := "" +
		"if (!electron_1.app.commandLine.hasSwitch('remote-debugging-port')) {\n" +
		"    electron_1.app.commandLine.appendSwitch('remote-debugging-port', '0');\n" +
		"}\n" +
		"electron_1.app.commandLine.appendSwitch('disable-blink-features', 'AutomationControlled');\n"
	locale := &Locale{Language: "zh-CN", Name: "简体中文", ChromiumLang: "zh-CN", Text: map[string]string{"Hello": "你好"}}
	out, err := applyLanguageSwitch(source, locale)
	if err != nil {
		t.Fatal(err)
	}
	if !langSwitchRe.MatchString(out) {
		t.Fatalf("language switch was not inserted:\n%s", out)
	}
	if !debugPortRe.MatchString(out) {
		t.Fatal("debugging-port anchor was rewritten instead of kept")
	}
	main, err := patchMain(out, locale, "/* antigravity-i18n:begin */\nfunction installLocaleDialogs(){}\n/* antigravity-i18n:end */")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(main, "installLocaleDialogs") {
		t.Fatal("dialog fragment was not applied")
	}
}
