package core

import (
	"fmt"
	"os"
	"testing"

	"github.com/dop251/goja"
)

// Exercise the shipped JavaScript engine, including first-match rule selection,
// so a generic device prompt cannot hide the more specific page descriptions.
func TestEngineDeviceConnectionPrompts(t *testing.T) {
	assets := os.DirFS("../..")
	store, err := NewLocaleStore(assets, assets, "assets/locales", "assets/patches")
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		name   string
		prompt string
	}{
		{"origin", "example.com wants to connect to %s"},
		{"embedded-origin", "An embedded page at example.com wants to connect to %s"},
		{"embedded-here", "An embedded page on this page wants to connect to %s"},
		{"this-page", "This page wants to connect to %s"},
	}
	devices := []string{"a HID device", "a USB device", "a serial port"}
	locales := []struct {
		code    string
		devices [3]string
		prompts [4]string
	}{
		{"de", [3]string{"einem HID-Gerät", "einem USB-Gerät", "einem seriellen Anschluss"}, [4]string{
			"example.com möchte eine Verbindung zu %s herstellen",
			"Eine eingebettete Seite unter example.com möchte eine Verbindung zu %s herstellen",
			"Eine eingebettete Seite auf dieser Seite möchte eine Verbindung zu %s herstellen",
			"Diese Seite möchte eine Verbindung zu %s herstellen",
		}},
		{"es", [3]string{"un dispositivo HID", "un dispositivo USB", "un puerto serie"}, [4]string{
			"example.com quiere conectarse a %s",
			"Una página incrustada en example.com quiere conectarse a %s",
			"Una página incrustada en esta página quiere conectarse a %s",
			"Esta página quiere conectarse a %s",
		}},
		{"fr", [3]string{"un périphérique HID", "un périphérique USB", "un port série"}, [4]string{
			"example.com veut se connecter à %s",
			"Une page intégrée sur example.com veut se connecter à %s",
			"Une page intégrée sur cette page veut se connecter à %s",
			"Cette page veut se connecter à %s",
		}},
		{"ja", [3]string{"HID デバイス", "USB デバイス", "シリアルポート"}, [4]string{
			"example.com が%sへの接続を求めています",
			"example.com の埋め込みページが%sへの接続を求めています",
			"このページ上の埋め込みページが%sへの接続を求めています",
			"このページが%sへの接続を求めています",
		}},
		{"ko", [3]string{"HID 기기", "USB 기기", "직렬 포트"}, [4]string{
			"example.com이(가) %s에 연결하려고 합니다",
			"example.com의 삽입된 페이지가 %s에 연결하려고 합니다",
			"이 페이지의 삽입된 페이지가 %s에 연결하려고 합니다",
			"이 페이지가 %s에 연결하려고 합니다",
		}},
		{"pt-BR", [3]string{"um dispositivo HID", "um dispositivo USB", "uma porta serial"}, [4]string{
			"example.com quer se conectar a %s",
			"Uma página incorporada em example.com quer se conectar a %s",
			"Uma página incorporada nesta página quer se conectar a %s",
			"Esta página quer se conectar a %s",
		}},
		{"ru", [3]string{"HID-устройству", "USB-устройству", "последовательному порту"}, [4]string{
			"example.com запрашивает подключение к %s",
			"Встроенная страница на example.com запрашивает подключение к %s",
			"Встроенная страница на этой странице запрашивает подключение к %s",
			"Эта страница запрашивает подключение к %s",
		}},
		{"zh-CN", [3]string{"HID 设备", "USB 设备", "串行端口"}, [4]string{
			"example.com 想要连接到%s",
			"example.com 上的嵌入页面想要连接到%s",
			"此页面上的嵌入页面想要连接到%s",
			"此页面想要连接到%s",
		}},
		{"zh-Hant", [3]string{"HID 裝置", "USB 裝置", "序列埠"}, [4]string{
			"example.com 想要連線到%s",
			"example.com 上的嵌入頁面想要連線到%s",
			"此頁面上的嵌入頁面想要連線到%s",
			"此頁面想要連線到%s",
		}},
	}
	for _, locale := range locales {
		t.Run(locale.code, func(t *testing.T) {
			loc, err := store.LoadLocale(locale.code)
			if err != nil {
				t.Fatal(err)
			}
			engine, err := store.buildPreloadFragment(loc)
			if err != nil {
				t.Fatal(err)
			}
			runtime := goja.New()
			// Defer DOM installation; these checks call the real string translator.
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
			for i, prompt := range cases {
				for j, device := range devices {
					t.Run(prompt.name+"/"+device, func(t *testing.T) {
						input := fmt.Sprintf(prompt.prompt, device)
						want := fmt.Sprintf(locale.prompts[i], locale.devices[j])
						got, err := translate(goja.Undefined(), runtime.ToValue(input))
						if err != nil {
							t.Fatal(err)
						}
						if got.String() != want {
							t.Fatalf("translateString(%q) = %q, want %q", input, got.String(), want)
						}
					})
				}
			}
		})
	}
}
