# <img src="./frontend/dist/appicon.png" width="32" height="32" alt=""> antigravity-i18n

為 Windows 版 Google Antigravity 提供九種介面語言套件，翻譯桌面應用程式裡的設定、對話區按鈕、選單和結束提示。也可以在輸入框旁查看剩餘的上下文和模型的輸出速度。設定生效後即可關閉本工具。

[下載 Windows x64 版](https://github.com/wadewu-ml/antigravity-i18n/releases/latest/download/antigravity-i18n.exe) · [版本記錄](https://github.com/wadewu-ml/antigravity-i18n/releases)

適用於 **Antigravity App 2.33.0**。

[简体中文](./README.zh-CN.md) · [English](./README.md) · **繁體中文** · [日本語](./README.ja.md) · [한국어](./README.ko.md) · [Deutsch](./README.de.md) · [Español](./README.es.md) · [Français](./README.fr.md) · [Português (Brasil)](./README.pt-BR.md) · [Русский](./README.ru.md)

<a href="./docs/images/antigravity-zh-cn.png"><img src="./docs/images/translated-window.webp" width="680" alt="簡體中文介面，頂部是中文選單，輸入框內可見模型選擇器與指標圓環"></a>

## 切換介面語言

可選擇簡體中文、繁體中文、日文、韓文、德文、西班牙文、法文、巴西葡萄牙文和俄文。工具本身的按鈕使用簡體中文。

1. 按兩下 `antigravity-i18n.exe`。如果工具找不到正確的安裝位置，點選「更改位置」，選擇含有 `Antigravity.exe` 的資料夾。
2. 在「界面语言」選擇語言。想查看指標就開啟「会话指标」。若要搭配原版英文介面，語言選 **English** 即可。
3. 儲存工作並結束 Antigravity，再點選「应用更改」。Antigravity 會以新設定重新開啟。

結束提示中的「强制退出并重试」會中斷進行中的工作階段，尚未儲存的變更可能會遺失。

<a href="./docs/images/language-settings.webp"><img src="./docs/images/language-settings.webp" width="640" alt="工具中已選取簡體中文，並開啟了會話指標"></a>


## 查看會話指標

對話有了模型呼叫紀錄後，輸入框內的模型選擇器右側就會出現圓環，顯示上下文已用比例的估計值。將游標移到圓環上，可查看剩餘 token 數、輸出速度、快取命中率，以及等待第一個 token 的時間。點一下固定詳細資料，再點一下即可收起。

<a href="./docs/images/session-metrics-zh-cn.png"><img src="./docs/images/session-metrics-detail.webp" width="560" alt="圓環上方的簡體中文面板顯示上下文用量、剩餘 token 和輸出速度"></a>

指標取自本機的會話紀錄。分割畫面、側問和子代理面板各自顯示對應會話最近一次有資料的模型呼叫。輸出量包含模型記錄的思考 token。`—` 表示該項資料暫時無法取得。

## 還原與更新

若要恢復原版英文介面並移除指標，請先結束 Antigravity，在工具中點選「还原官方版本」，再點選「还原」確認。首次套用時，工具會自動備份原始檔案。請保留安裝目錄下的 `resources/app.asar.clean-backup` 供還原使用。

如果 Antigravity 更新後翻譯或指標消失，請先對照本頁頂端列出的適用版本，再選擇語言和指標並套用。要再次套用目前生效的設定，請點選「重新应用」。

## 命令列用法

在 `antigravity-i18n.exe` 所在目錄開啟 PowerShell。

| 操作 | 指令 |
| --- | --- |
| 列出語言代碼 | `.\antigravity-i18n.exe locales` |
| 套用繁體中文 | `.\antigravity-i18n.exe apply --locale zh-Hant` |
| 使用官方英文並開啟指標 | `.\antigravity-i18n.exe apply --locale en --session-metrics` |
| 查看目前設定 | `.\antigravity-i18n.exe status` |
| 還原官方版本 | `.\antigravity-i18n.exe restore` |
| 查看所有選項 | `.\antigravity-i18n.exe --help` |

## 從原始碼建置

在 Windows 上安裝 PowerShell 7、Go 1.26.3 或更新版本，以及 Wails v2 CLI。在儲存庫根目錄執行 `.\build.ps1 -Target gui`，產物位於 `build/bin/antigravity-i18n.exe`。

## 啟動問題

如果啟動工具時提示缺少 WebView2，請安裝 [Microsoft Edge WebView2 執行階段](https://developer.microsoft.com/en-us/microsoft-edge/webview2/#download)。

## 免責聲明

本專案為開源社區維護的非官方本地化工具，與 Google 或 Antigravity 官方無關。

[MIT 授權條款](./LICENSE)
