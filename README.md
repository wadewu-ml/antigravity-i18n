# <img src="./frontend/dist/appicon.png" width="32" height="32" alt=""> antigravity-i18n

Nine language packs translate settings, chat controls, menus, and quit prompts in the Google Antigravity desktop app for Windows. You can also see how much context is left and how fast the model responded, right in the chat input. Once your settings are applied, you can close the tool.

[Download for Windows x64](https://github.com/wadewu-ml/antigravity-i18n/releases/latest/download/antigravity-i18n.exe) · [Releases](https://github.com/wadewu-ml/antigravity-i18n/releases)

Tested with **Antigravity App 2.35.1**.

[简体中文](./README.zh-CN.md) · **English** · [繁體中文](./README.zh-Hant.md) · [日本語](./README.ja.md) · [한국어](./README.ko.md) · [Deutsch](./README.de.md) · [Español](./README.es.md) · [Français](./README.fr.md) · [Português (Brasil)](./README.pt-BR.md) · [Русский](./README.ru.md)

<a href="./docs/images/antigravity-zh-cn.png"><img src="./docs/images/translated-window.webp" width="680" alt="Antigravity in Simplified Chinese, with translated menus and the metrics ring beside the model selector"></a>

## Change the interface language

The packs cover Simplified Chinese, Traditional Chinese, Japanese, Korean, German, Spanish, French, Brazilian Portuguese, and Russian. The setup tool uses Simplified Chinese.

1. Open `antigravity-i18n.exe`. If the installation folder is wrong, use the button beside the path to select the folder containing `Antigravity.exe`.
2. Choose your language. To see session metrics, turn on the switch below the language list. You can also use it with English.
3. Save your work, quit Antigravity, and click Apply changes (应用更改) at the bottom right. Antigravity will reopen.

If prompted to quit, save your work first. Forcing Antigravity to close interrupts active sessions and can discard unsaved edits.

To enable the optional slider, turn on the switch below Session metrics in the setup tool and apply your changes. Then open the model picker. For models with selectable thinking effort, drag the slider above View Usage and release it to apply the level.

<a href="./docs/images/language-settings.webp"><img src="./docs/images/language-settings.webp" width="640" alt="Settings window with Simplified Chinese selected and session metrics enabled"></a>

The green button shown here reads 重新应用 (Reapply).

## Read session metrics

Hover over the ring beside the model selector to see each model group’s remaining quota by period. Hover over a percentage for the local reset time. You can check quota before sending your first message. Once the chat records a model call, the ring shows estimated context use and the panel shows tokens used and total capacity, output speed, cache hit rate, and time to first token. Click the ring to pin or dismiss the panel.

<a href="./docs/images/session-metrics-zh-cn.png"><img src="./docs/images/session-metrics-detail.webp" width="560" alt="The open panel above the ring shows context usage, remaining tokens, and output speed in Simplified Chinese"></a>

In split views, side chats, and subagent panels, each panel shows the latest model call with data in its own local session record. The output count includes any thinking tokens the model reports. A dash (`—`) means that a value is unavailable.

## Restore or reapply

To return to the original English interface and remove the indicator and effort slider, quit Antigravity, open the menu at the top right of the tool, click Restore official version (还原官方版本), then confirm. The first time you apply settings, the tool saves the original file in the Antigravity installation folder. Keep `resources/app.asar.clean-backup` there for restoration.

If an Antigravity update removes the translation or indicator, check the supported version at the top of this page, then select and apply your settings again. To reapply settings that are already active, click Reapply (重新应用).

To update the tool, open the menu at the top right of the window and click Check for updates (检查更新). It checks GitHub when you click. Confirm to download the new version and reopen the tool from the same EXE path. To apply the updated translations, click Reapply.

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

1. This is an unofficial, community-maintained project and is not affiliated with, endorsed by, or authorized by Google or Antigravity.
2. Google Antigravity and related trademarks belong to their respective owners.
3. This tool is provided for personal and educational use. Please ensure compliance with relevant software terms of service.

[MIT License](./LICENSE)
