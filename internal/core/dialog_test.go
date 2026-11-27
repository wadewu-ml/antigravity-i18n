package core

import (
	"testing"

	"github.com/dop251/goja"
)

func TestNativeDialogTranslationPreservesOptions(t *testing.T) {
	store := newTestService(t).store
	loc, err := store.LoadLocale(DefaultLocale)
	if err != nil {
		t.Fatal(err)
	}
	fragment, err := store.buildDialogFragment(loc)
	if err != nil {
		t.Fatal(err)
	}
	js := goja.New()
	_, err = js.RunString(`
const dialog = {showMessageBox(...args) { return args; }, showMessageBoxSync(...args) { return args; }};
const electron_1 = {dialog};
const win = {};
` + fragment + `
const quit = dialog.showMessageBox(win, {
    title: 'Confirm Quit', message: 'Are you sure you want to quit?', buttons: ['Cancel', 'Quit'], cancelId: 0,
});
const update = dialog.showMessageBoxSync({
    title: 'Check for Updates', message: 'No updates available', buttons: ['OK'], defaultId: 0,
});
const wsl = dialog.showMessageBox({
    type: 'warning', title: 'WSL distro not found', message: 'The WSL distro "My Distro $& {1}" is no longer installed.',
    detail: 'Antigravity opened on Windows instead.',
});
const unrelated = {title: 'Check for Updates', message: 'Network error', buttons: ['OK']};
`)
	if err != nil {
		t.Fatal(err)
	}
	for expr, want := range map[string]string{
		"quit[0] === win":           "true",
		"quit[1].buttons.join('/')": "取消/退出",
		"quit[1].cancelId":          "0",
		"update[0].title":           "检查更新",
		"update[0].message":         "暂无可用更新",
		"update[0].buttons[0]":      "确定",
		"update[0].defaultId":       "0",
		"wsl[0].type":               "warning",
		"wsl[0].title":              "找不到 WSL 发行版",
		"wsl[0].message":            "WSL 发行版「My Distro $& {1}」已卸载。",
		"wsl[0].detail":             "Antigravity 已改在 Windows 中打开。",
		"dialog.showMessageBox(unrelated)[0] === unrelated": "true",
	} {
		got, err := js.RunString(expr)
		if err != nil || got.String() != want {
			t.Errorf("%s: got %v, %v, want %q", expr, got, err, want)
		}
	}
}
