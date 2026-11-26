<div align="center">
  <img src="./frontend/dist/appicon.png" width="64" height="64" alt="Regravity icon">
  <h1>Regravity</h1>
  <p>Localization packs, session metrics & model quotas, and a thinking-effort slider for Google Antigravity on Windows.</p>
  <p>
    <a href="https://github.com/wadewu-ml/regravity/releases/download/v4.2.3/antigravity-i18n.exe"><strong>Download for Windows x64</strong></a>
    ·
    <a href="https://github.com/wadewu-ml/regravity/releases">Releases</a>
    ·
    Tested with <strong>Antigravity App 2.35.1</strong>
  </p>
  <p>
    <a href="./README.zh-CN.md">简体中文</a> · <strong>English</strong> · <a href="./README.zh-Hant.md">繁體中文</a> · <a href="./README.ja.md">日本語</a> · <a href="./README.ko.md">한국어</a> · <a href="./README.de.md">Deutsch</a> · <a href="./README.es.md">Español</a> · <a href="./README.fr.md">Français</a> · <a href="./README.pt-BR.md">Português (Brasil)</a> · <a href="./README.ru.md">Русский</a>
  </p>
  <p>Regravity is the new name of antigravity-i18n. The download above is v4.2.3. When the first Regravity release appears, download <code>regravity.exe</code> manually. Your existing patches and update preference carry over. The v4.2.3 update checker can report a version error after the repository rename. Use the Releases link above.</p>
</div>

<p align="center">
  <a href="./docs/images/language-settings.webp">
    <img src="./docs/images/language-settings.webp" width="560" alt="Regravity setup window showing Antigravity v2.35.2, Simplified Chinese selected, and both Session metrics and Thinking-effort slider enabled">
  </a>
</p>

## Features

- **9 interface languages in one portable EXE** — Translates settings, chat controls, model menus, native menus, and quit prompts into Simplified Chinese, Traditional Chinese, Japanese, Korean, German, Spanish, French, Brazilian Portuguese, and Russian without touching your chats or input.
- **Optional session metrics & model quotas** — Adds a ring beside the model selector to inspect remaining model quotas, local reset times, context used, average output speed, cache hit rate, and time to first token.
- **Optional thinking-effort slider** — Adds an inline capsule slider inside the model picker for models with selectable thinking effort.
- **Independent & reversible** — Both optional switches work independently and with the built-in English UI. Once applied, you can close the tool and restore the official version byte-for-byte using a clean backup that matches the installed Antigravity version.

## Quick start

> [!NOTE]
> The setup tool UI is currently in Simplified Chinese. Original labels are included in parentheses below to match the screenshot above.

### Apply settings

1. Run `antigravity-i18n.exe`. If the detected installation path is wrong, click **Change location** (`更改位置`) and select the folder containing `Antigravity.exe`.
2. Pick an interface language under **Interface language** (`界面语言`), or select **English** to keep the original UI while using the optional switches.
3. Toggle **Session metrics** (`会话指标`) and/or **Thinking-effort slider** (`思考强度滑块`) below the language list.
4. Save your work, quit Antigravity, and click **Apply changes** (`应用更改`) at the bottom right. Antigravity will reopen with your settings (or click **Open Antigravity** / `打开 Antigravity`), and you can close the tool.

> [!WARNING]
> If prompted that Antigravity is still running, save your work and quit normally. **Force quit and retry** (`强制退出并重试`) interrupts active sessions and can discard unsaved edits.

### Reapply or restore

- **Reapply current settings** — Click **Reapply** (`重新应用`) at the bottom right.
- **Restore official version** — Quit Antigravity, open the top-right menu (`···`), click **Restore official version** (`还原官方版本`), and confirm. This restores the original English UI and removes the metrics ring and effort slider.

> [!IMPORTANT]
> On first apply, the tool saves a clean backup at `resources/app.asar.clean-backup` inside the Antigravity installation folder. Keep that file so you can restore later.

## Optional controls

### Session metrics & model quotas

- **Quotas before the first message** — Hover over the ring beside the model selector to see remaining quota per model group and period (for example weekly and 5-hour windows). Hover over a percentage to see the local reset time.
- **Live metrics after a model call** — Once a chat records a model call, the ring reflects estimated context usage and the upper section fills in **Context used** (tokens used, total capacity, and percentage), **Average output speed** (including reported thinking tokens), **Cache hit rate**, and **Time to first token**. Click the ring to pin or dismiss the panel.
- **Per-panel tracking** — Split views, side chats, and subagent panels each read the latest model call with data from their own local session record. A dash (`—`) indicates a value is unavailable.

<p align="center">
  <a href="./docs/images/session-quota-before-first-call.webp">
    <img src="./docs/images/session-quota-before-first-call.webp" width="440" alt="Session metrics and model quota remaining panel above the ring beside the model selector">
  </a>
</p>

### Thinking-effort slider

Turn on **Thinking-effort slider** (`思考强度滑块`) and apply your changes. In Antigravity's model picker, models with selectable thinking effort show an inline slider above **View Usage** — drag across detents to preview a level and release to apply it.

<p align="center">
  <a href="./docs/images/thinking-effort-high.webp">
    <img src="./docs/images/thinking-effort-high.webp" height="340" alt="Model picker with thinking effort set to High">
  </a>
  &nbsp;&nbsp;
  <a href="./docs/images/thinking-effort-max.webp">
    <img src="./docs/images/thinking-effort-max.webp" height="340" alt="Model picker with thinking effort set to Max">
  </a>
</p>

## Updates

- **After an Antigravity update** — If translations or controls disappear, verify the tested version at the top of this page, then reopen the tool and apply your settings again.
- **Tool auto-updates** — At startup, the tool uses the previous check result to decide whether to contact GitHub. After a successful check, it waits at least 24 hours. A blue dot on the top-right menu (`···`) signals a new release; open the menu to read release notes and download and restart from the same EXE path. Automatic checks can be turned off in **About** (`关于`), and manual checks remain available via **Check for updates** (`检查更新`). After updating, click **Reapply** (`重新应用`) to push updated translations into Antigravity.

## PowerShell CLI

Run these commands from PowerShell in the folder containing `antigravity-i18n.exe`:

| Task | Command |
| --- | --- |
| List language codes | `.\antigravity-i18n.exe locales` |
| Apply Simplified Chinese | `.\antigravity-i18n.exe apply --locale zh-CN` |
| Keep English with session metrics | `.\antigravity-i18n.exe apply --locale en --session-metrics` |
| Keep English with thinking-effort slider | `.\antigravity-i18n.exe apply --locale en --effort-slider` |
| Check current status | `.\antigravity-i18n.exe status` |
| Restore official version | `.\antigravity-i18n.exe restore` |
| Show all options | `.\antigravity-i18n.exe --help` |

## Build from source

Requires Windows, PowerShell 7, Go 1.26.3+, and the Wails v2 CLI:

```powershell
.\build.ps1 -Target gui
```

The built executable is written to `build/bin/regravity.exe`.

## Troubleshooting

- **Missing WebView2 prompt on startup** — Install the [Microsoft Edge WebView2 Runtime](https://developer.microsoft.com/en-us/microsoft-edge/webview2/#download) and relaunch `antigravity-i18n.exe`.

## Disclaimer

1. This is an unofficial, community-maintained project and is not affiliated with, endorsed by, or authorized by Google or Antigravity.
2. Google Antigravity and related trademarks belong to their respective owners.
3. This tool is provided for personal and educational use. Please ensure compliance with relevant software terms of service.

[MIT License](./LICENSE)
