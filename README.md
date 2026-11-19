# <img src="./frontend/dist/appicon.png" width="32" height="32" alt=""> antigravity-i18n

Nine language packs translate settings, chat controls, menus, and quit prompts in the Google Antigravity desktop app for Windows. You can also see how much context is left and how fast the model responded, right in the chat input. Once your settings are applied, you can close the tool.

[Download for Windows x64](https://github.com/wadewu-ml/antigravity-i18n/releases/latest/download/antigravity-i18n.exe) · [Releases](https://github.com/wadewu-ml/antigravity-i18n/releases)

Tested with **Antigravity App 2.33.0**.

[简体中文](./README.zh-CN.md) · **English** · [繁體中文](./README.zh-Hant.md) · [日本語](./README.ja.md) · [한국어](./README.ko.md) · [Deutsch](./README.de.md) · [Español](./README.es.md) · [Français](./README.fr.md) · [Português (Brasil)](./README.pt-BR.md) · [Русский](./README.ru.md)

<a href="./docs/images/antigravity-zh-cn.png"><img src="./docs/images/translated-window.webp" width="680" alt="Antigravity in Simplified Chinese, with translated menus and the metrics ring beside the model selector"></a>

## Change the interface language

The packs cover Simplified Chinese, Traditional Chinese, Japanese, Korean, German, Spanish, French, Brazilian Portuguese, and Russian. The setup tool uses Simplified Chinese.

1. Open `antigravity-i18n.exe`. If the tool cannot find the right installation, click **更改位置** and select the folder containing `Antigravity.exe`.
2. Choose a language under **界面语言**. Turn on **会话指标** if you want session metrics. For metrics with the original interface, select **English**.
3. Save your work and quit Antigravity, then click **应用更改** to apply. Antigravity reopens with your new settings.

If prompted, **强制退出并重试** forces Antigravity to quit and retries the operation. This interrupts active sessions and can discard unsaved edits.

<a href="./docs/images/language-settings.webp"><img src="./docs/images/language-settings.webp" width="640" alt="Settings window with Simplified Chinese selected and session metrics enabled"></a>


## Read session metrics

Once a chat has a recorded model call, a ring appears to the right of the model selector in the chat input. Its fill shows the estimated share of the context window used. Hover to see the remaining tokens, output speed, cache hit rate, and time to first token. Click the ring to pin the details and click again to dismiss them.

<a href="./docs/images/session-metrics-zh-cn.png"><img src="./docs/images/session-metrics-detail.webp" width="560" alt="The open panel above the ring shows context usage, remaining tokens, and output speed in Simplified Chinese"></a>

In split views, side chats, and subagent panels, each panel shows the latest model call with data in its own local session record. The output count includes any thinking tokens the model reports. A dash (`—`) means that a value is unavailable.

## Restore or reapply

To return to the original English interface and remove the indicator, quit Antigravity, click **还原官方版本** in the tool, then confirm with **还原**. The first time you apply settings, the tool saves the original file in the Antigravity installation folder. Keep `resources/app.asar.clean-backup` there for restoration.

If an Antigravity update removes the translation or indicator, check the supported version at the top of this page, then select and apply your settings again. To reapply settings that are already active, click **重新应用**.

## PowerShell commands

Open PowerShell in the folder containing `antigravity-i18n.exe`.

| Task | Command |
| --- | --- |
| List language codes | `.\antigravity-i18n.exe locales` |
| Apply Simplified Chinese | `.\antigravity-i18n.exe apply --locale zh-CN` |
| Use built-in English with metrics | `.\antigravity-i18n.exe apply --locale en --session-metrics` |
| Check current settings | `.\antigravity-i18n.exe status` |
| Restore the official version | `.\antigravity-i18n.exe restore` |
| See all options | `.\antigravity-i18n.exe --help` |

## Build from source

On Windows, install PowerShell 7, Go 1.26.3 or later, and the Wails v2 CLI. Run `.\build.ps1 -Target gui` from the repository root to build `build/bin/antigravity-i18n.exe`.

## Startup troubleshooting

If the tool asks for WebView2 at startup, install the [Microsoft Edge WebView2 Runtime](https://developer.microsoft.com/en-us/microsoft-edge/webview2/#download).

## Disclaimer

This is an unofficial, community-maintained project and is not affiliated with, endorsed by, or authorized by Google or Antigravity.

[MIT License](./LICENSE)
