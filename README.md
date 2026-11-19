# antigravity-i18n

Portable Windows tool for UI localization and optional session metrics in the Google Antigravity desktop client. [Google Antigravity](https://antigravity.google/)

[简体中文](./README.zh-CN.md) | English | [繁體中文](./README.zh-Hant.md) | [日本語](./README.ja.md) | [한국어](./README.ko.md) | [Deutsch](./README.de.md) | [Español](./README.es.md) | [Français](./README.fr.md) | [Português (Brasil)](./README.pt-BR.md) | [Русский](./README.ru.md)

---

## Compatibility

v4.0.0 supports Antigravity App **2.33.0**. All nine packs cover the new browser data controls, shortcut schemes, built-in slash command descriptions, SSH host controls, GitHub connection errors, side question actions, plugin bundle dialogs, and voice dictation status. Applying, switching languages, and restoring have been verified against the official 2.33.0 archive. [Antigravity changelog](https://antigravity.google/docs/changelog#rel-hub-2.33.0)

The [2.32.1 audit](./docs/translation-audit-2.32.1.md) describes the earlier updates to settings, onboarding, and related screens. The engine preserves spoken transcripts, call summaries, AI progress text, and user text in side questions. It also preserves links, controls, and values such as hosts and search terms in translated sentences. The Chinese packs keep the permission preset name Turbo. Dictionary alignment measures keys in the bundled packs. The audit lists surfaces and text outside its scope.

---

## Key Features

- **Portable Windows x64 executable**. Run the file directly. The graphical interface uses Microsoft Edge WebView2.
- **GUI & CLI Dual Mode**: Opens a clean desktop window by default, while embedding complete CLI subcommands (`apply` / `restore` / `status` / `locales`) for scripting and automation.
- **Byte-Exact Restoration**: Automatically creates a pristine backup (`app.asar.clean-backup`) before any changes. Restoring recovers an archive identical byte-for-byte (SHA-256 verified) to the official release.
- **Interface localization** covers the shell, native menus, and quit dialogs.
- **Optional session metrics** show context usage in a ring to the right of the model selector. Hover for remaining context, recent output speed, cache hit rate, and time to first token, or click to keep the details open. Antigravity reads the metrics from its local database in read-only mode.
- **Local patch processing** includes JavaScript syntax validation and transactional rollback on errors.

---

## Download & Usage

### Download

Download the latest `antigravity-i18n.exe` from the [Releases](../../releases) page and run it directly.

### 1. Graphical Interface (Recommended)

Double-click `antigravity-i18n.exe`:
1. The app automatically detects the Antigravity installation folder (or click "Change Location" to specify manually).
2. Select a language and set the “会话指标” switch independently. The metrics also work with the built-in English interface.
3. Click “应用更改” to apply and restart Antigravity. Use “还原官方版本” to restore the official English archive and remove the indicator.

Split views, side questions, and subagent panes show metrics for their own conversation. A dash indicates unavailable data. You can close the installer after applying the patch.

When apply or restore needs to close Antigravity, the tool asks the selected installation to exit normally. If it is still running after 30 seconds, you can close it manually and try again or click “强制退出并重试” (Force quit and retry).

### 2. Command Line Interface

Open PowerShell in the folder containing `antigravity-i18n.exe`.

```powershell
# Check current localization status and backups
.\antigravity-i18n.exe status

# Install a language pack (Simplified Chinese by default)
.\antigravity-i18n.exe apply --locale zh-CN

# Restore the official untranslated English version
.\antigravity-i18n.exe restore

# List all bundled language packs with coverage
.\antigravity-i18n.exe locales
```

Common options:
- `--app-dir <path>`: Manually specify the Antigravity installation directory (or set via `ANTIGRAVITY_APP_DIR`).
- `--locale <code>`: Target language code (e.g., `zh-CN`, `zh-Hant`, `ja`, `ko`, `de`, `es`, `fr`).
- Add `--session-metrics` to `apply` to enable the indicator, or `--session-metrics=false` to disable it. `apply --locale en --session-metrics` uses built-in English with the indicator. For bundled language packs, omitting the flag uses the setting in the installed archive.
- `--no-restart`: Do not relaunch Antigravity after patching.
- `--force` immediately closes the selected installation. Unsaved work may be lost.
- `--timeout <ms>` sets how long to wait for Antigravity to exit before stopping the operation. The default is `30000`.
- `--yes`: Suppress the shutdown confirmation prompt.

Apply and restore allow one operation at a time per installation. If the tool reports that another operation is running, retry after it finishes. Use `status` to check the selected installation and its backup. An unreadable archive or invalid application identity produces an error.

---

## Bundled Languages

All language packs are embedded directly into the executable:

Coverage below measures entry parity with the bundled reference dictionary, not every English string in the client. Spoken content, user content, and sentence fragments without a complete match are kept in their original language.

| Code | Language | Coverage | Notes |
| :--- | :--- | :--- | :--- |
| `zh-CN` | Simplified Chinese (简体中文) | 100% | Full coverage of shell UI, native menus, and dialogs |
| `zh-Hant` | Traditional Chinese (繁體中文) | 100% | Full coverage with traditional terminology |
| `ja` | Japanese (日本語) | 100% | Full coverage with verified technical terminology |
| `ko` | Korean (한국어) | 100% | Full coverage with verified technical terminology |
| `de` | German (Deutsch) | 100% | Full coverage with verified technical terminology |
| `es` | Spanish (Español) | 100% | Full coverage with verified technical terminology |
| `fr` | French (Français) | 100% | Full coverage with verified technical terminology |
| `pt-BR` | Portuguese (Brasil) | 100% | Full coverage with verified technical terminology |
| `ru` | Russian (Русский) | 100% | Full coverage with verified technical terminology |
| `en` | English | - | Official untouched build (select to restore) |

---

## FAQ & Mechanics

### 1. What happens after an official Antigravity update?
Official updates will overwrite `app.asar`. Simply re-open this tool and click "Apply" again. The tool will automatically refresh its clean backup using the new official release.

### 2. Language switching and restoration
You can switch directly between any language; the tool seamlessly replaces active patches without requiring a manual restore first. Clicking "Restore Official English" returns the client to an untouched state at any time.

### 3. Backup file management
- `app.asar.clean-backup` stores the original archive used for restoration. When you apply or restore an unpatched installation, the tool checks the backup against the installed archive by SHA-256 and refreshes it if the hashes differ, including after a same-version reinstall. **Keep this file.**
- `app.asar.bak-*` contains timestamped copies that can help recover a missing or unusable clean backup. Once the client works and `status` shows a valid clean backup, you can remove older snapshots to free disk space.

### 4. Why do some words or sentences remain in English?
Translation matches complete text rather than individual words in an unknown sentence. Fragments split by inline markup are also left intact to avoid partially translated sentences. Independent buttons, menu items, and accessibility labels can still be translated.

Voice mode preserves user speech, word-highlighted replies, and transcript text, while translating voice controls and work status. The Chinese packs retain the permission preset name `Turbo`, whose behavior concerns automatic execution and access permissions rather than speed alone.

---

## Building from Source

To compile from source:

**Prerequisites**:
- Go 1.26.3+
- [Wails v2](https://wails.io/) CLI
- PowerShell 7

```powershell
# Clone the repository
git clone https://github.com/wadewu-ml/antigravity-i18n.git
cd antigravity-i18n

# Build both GUI and CLI binaries
.\build.ps1

# Uncomment the next line to include an official-archive round trip.
# $env:ANTIGRAVITY_ASAR = "$env:LOCALAPPDATA\Programs\antigravity\resources\app.asar.clean-backup"
go test ./...
```

The DOM tests use a separate headless Chrome or Edge instance. Set `ANTIGRAVITY_TEST_BROWSER` to specify another Chromium executable. If no browser is found, the DOM tests are skipped.

---

## Disclaimer

1. This is an unofficial, community-maintained project and is not affiliated with, endorsed by, or authorized by Google or Antigravity.
2. Google Antigravity and related trademarks belong to their respective owners.
3. This tool is provided for personal and educational use. Please ensure compliance with relevant software terms of service.

## License

[MIT License](./LICENSE)
