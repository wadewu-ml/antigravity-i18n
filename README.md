<div align="center">
  <img src="./frontend/dist/appicon.png" width="68" height="68" alt="Regravity icon">
  <h1>Regravity</h1>
  <p><em>Antigravity takes off; Regravity brings it back to your language and control.</em></p>
  <p>Portable Windows companion for <strong>Google Antigravity</strong> — 9 interface languages, session metrics &amp; model quotas, and a thinking-effort slider.</p>
  <p>
    <a href="https://github.com/wadewu-ml/regravity/releases/download/v4.2.3/antigravity-i18n.exe"><img src="https://img.shields.io/badge/Download-v4.2.3_(Windows_x64)-0F9D58?style=flat&logo=windows&logoColor=white" alt="Download v4.2.3 for Windows x64"></a>
    <a href="https://github.com/wadewu-ml/regravity/releases"><img src="https://img.shields.io/badge/Releases-GitHub-18181B?style=flat&logo=github" alt="GitHub Releases"></a>
    <img src="https://img.shields.io/badge/Antigravity-2.35.1_tested-18181B?style=flat" alt="Tested with Antigravity App 2.35.1">
    <a href="./LICENSE"><img src="https://img.shields.io/badge/License-MIT-18181B?style=flat" alt="MIT License"></a>
  </p>
  <p><sub>Current Windows file · <code>antigravity-i18n.exe</code></sub></p>
  <p>
    <sub>
      <a href="./README.zh-CN.md">简体中文</a> · <strong>English</strong> · <a href="./README.zh-Hant.md">繁體中文</a> · <a href="./README.ja.md">日本語</a> · <a href="./README.ko.md">한국어</a> · <a href="./README.de.md">Deutsch</a> · <a href="./README.es.md">Español</a> · <a href="./README.fr.md">Français</a> · <a href="./README.pt-BR.md">Português (Brasil)</a> · <a href="./README.ru.md">Русский</a>
    </sub>
  </p>
</div>

<br>

<p align="center">
  <a href="./docs/images/language-settings.webp">
    <img src="./docs/images/language-settings.webp" width="580" alt="Regravity setup window showing Antigravity v2.35.2, Simplified Chinese selected, and both Session metrics and Thinking-effort slider enabled">
  </a>
  <br>
  <sub>Regravity development build with Antigravity 2.35.2.</sub>
</p>

## What's inside

- **9 interface languages, zero chat interference** — Translates settings, chat controls, model menus, native menus, and quit prompts into `zh-CN`, `zh-Hant`, `ja`, `ko`, `de`, `es`, `fr`, `pt-BR`, and `ru` while leaving your prompts and conversations untouched.
- **Session metrics & model quotas** *(optional)* — Adds an unobtrusive ring beside the model selector for remaining quotas, local reset times, context usage, output speed, cache hit rate, and time to first token.
- **Thinking-effort slider** *(optional)* — Embeds a capsule slider inside the model picker for models with selectable thinking effort.
- **Modular & byte-for-byte reversible** — Both optional switches work independently and with the built-in English UI. Close the tool once applied, and restore the official build from a matching clean backup whenever you want.

## Quick start

> [!TIP]
> The setup window is currently in Simplified Chinese. Original labels appear in parentheses below so you can match them directly against the screenshot above.

### Apply settings

1. Run `antigravity-i18n.exe`. If the detected folder is wrong, click **Change location** (`更改位置`) and select the folder containing `Antigravity.exe`.
2. Choose a language under **Interface language** (`界面语言`), or keep **English** if you only want the optional controls.
3. Enable **Session metrics** (`会话指标`) and/or **Thinking-effort slider** (`思考强度滑块`) below the language list.
4. Save your work, quit Antigravity, and click **Apply changes** (`应用更改`) at the bottom right. Antigravity reopens with your settings (or launch it via **Open Antigravity** / `打开 Antigravity`), and you can close the tool.

> [!WARNING]
> If prompted that Antigravity is still running, save your work and quit normally. **Force quit and retry** (`强制退出并重试`) interrupts active sessions and can discard unsaved edits.

### Reapply or restore

- **Reapply** — Click **Reapply** (`重新应用`) at the bottom right to reapply active settings.
- **Restore official version** — Quit Antigravity, open the top-right menu (`···`), click **Restore official version** (`还原官方版本`), and confirm to return to the original English build without the ring or slider.

> [!IMPORTANT]
> On first apply, the tool saves an unmodified copy at `resources/app.asar.clean-backup` inside the Antigravity folder. Keep that file if you want to restore later.

## Optional controls

### Session metrics & model quotas

- **Before the first message** — Hover over the ring beside the model selector to view remaining quota per model group and window (such as weekly and 5-hour limits); hover over a percentage for the local reset time.
- **After a model call** — The ring tracks estimated context usage, and the upper panel fills in **Context used** (tokens, total capacity, and share), **Average output speed** (including reported thinking tokens), **Cache hit rate**, and **Time to first token**. Click the ring to pin or dismiss.
- **Per-panel accuracy** — Split views, side chats, and subagent panels each read the latest model call with data from their own local session record. A dash (`—`) means unavailable.

<p align="center">
  <a href="./docs/images/session-quota-before-first-call.webp">
    <img src="./docs/images/session-quota-before-first-call.webp" width="420" alt="Session metrics and model quota remaining panel above the ring beside the model selector">
  </a>
  <br>
  <sub>Live quota breakdown is available immediately; session metrics populate after the first model call.</sub>
</p>

### Thinking-effort slider

Enable **Thinking-effort slider** (`思考强度滑块`) and apply your changes. In Antigravity's model picker, supported models show an inline slider above **View Usage** — drag across detents to preview and release to apply.

<p align="center">
  <a href="./docs/images/thinking-effort-high.webp">
    <img src="./docs/images/thinking-effort-high.webp" height="340" alt="Model picker with thinking effort set to High">
  </a>
  &nbsp;&nbsp;
  <a href="./docs/images/thinking-effort-max.webp">
    <img src="./docs/images/thinking-effort-max.webp" height="340" alt="Model picker with thinking effort set to Max">
  </a>
  <br>
  <sub>Standard detent (High) and highlighted maximum level (Max) directly inside the model menu.</sub>
</p>

## Updates & CLI

### Updates

- **After an Antigravity update** — If translations or controls disappear, verify the tested client version at the top of this page, reopen the tool, and apply your settings again.
- **Tool auto-updates** — At startup, the tool uses the previous check result to decide whether to contact GitHub and waits at least 24 hours after a successful check. A blue dot on the top-right menu (`···`) indicates a new release; open it to inspect notes and update in place. Automatic checks can be toggled in **About** (`关于`), with **Check for updates** (`检查更新`) always available manually. Click **Reapply** (`重新应用`) after updating to apply new translations.

If the v4.2.3 updater reports a version error after the rename, download the EXE from [Releases](https://github.com/wadewu-ml/regravity/releases). When the first Regravity release is available, download `regravity.exe` manually. Your existing patches and automatic update preference carry over.

### PowerShell CLI

Run from PowerShell in the folder containing `antigravity-i18n.exe`:

| Task | Command |
| --- | --- |
| List language codes | `.\antigravity-i18n.exe locales` |
| Apply Simplified Chinese | `.\antigravity-i18n.exe apply --locale zh-CN` |
| Keep English + session metrics | `.\antigravity-i18n.exe apply --locale en --session-metrics` |
| Keep English + effort slider | `.\antigravity-i18n.exe apply --locale en --effort-slider` |
| Check current status | `.\antigravity-i18n.exe status` |
| Restore official version | `.\antigravity-i18n.exe restore` |
| Show all options | `.\antigravity-i18n.exe --help` |

### Build from source

Requires Windows, PowerShell 7, Go 1.26.3+, and the Wails v2 CLI:

```powershell
.\build.ps1 -Target gui
```

Outputs `build/bin/regravity.exe`.

## Troubleshooting

- **Missing WebView2 on startup** — Install the [Microsoft Edge WebView2 Runtime](https://developer.microsoft.com/en-us/microsoft-edge/webview2/#download) and relaunch `antigravity-i18n.exe`.

---

## Disclaimer

1. This is an unofficial, community-maintained project and is not affiliated with, endorsed by, or authorized by Google or Antigravity.
2. Google Antigravity and related trademarks belong to their respective owners.
3. This tool is provided for personal and educational use. Please ensure compliance with relevant software terms of service.

[MIT License](./LICENSE)
