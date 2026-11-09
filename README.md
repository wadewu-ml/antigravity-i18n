# antigravity-i18n

UI localization tool for the [Google Antigravity](https://antigravity.google/) desktop client.

[简体中文](./README.zh-CN.md) | English | [繁體中文](./README.zh-Hant.md) | [日本語](./README.ja.md) | [한국어](./README.ko.md) | [Deutsch](./README.de.md) | [Español](./README.es.md) | [Français](./README.fr.md) | [Português (Brasil)](./README.pt-BR.md) | [Русский](./README.ru.md)

---

A standalone executable bundling a modern graphical interface, a pure Go asar engine, and complete language packs. Runs out of the box with zero dependencies on Node.js or runtime libraries, and supports byte-exact restoration to the official build at any time.

---

## Compatibility

Version 3.0.6 targets Antigravity App 2.29.0. It keeps the 3.0.5 shell coverage and leaves spoken transcripts and word-highlighted replies untranslated, while voice controls still translate. The Chinese packs keep the permission preset name Turbo. Backup, restore, and language-pack injection are unchanged.

Checked against the 2.29.0 renderer: about 97% of public settings labels, titles, and descriptions are translated, including the new public strings from this release. Text that stays English is the Monaco and VS Code color catalog, editor and layout-library messages, KaTeX help, and Google-internal screens (ABFS, CitC, Cog, Chief of Staff, and the developer host simulator).

---

## Key Features

- **Single Self-Contained Binary**: ~12 MB standalone executable—no Node.js, npm, or Python runtime required.
- **GUI & CLI Dual Mode**: Opens a clean desktop window by default, while embedding complete CLI subcommands (`apply` / `restore` / `status` / `locales`) for scripting and automation.
- **Byte-Exact Restoration**: Automatically creates a pristine backup (`app.asar.clean-backup`) before any changes. Restoring recovers an archive identical byte-for-byte (SHA-256 verified) to the official release.
- **Non-Invasive Design**: Localizes only the shell UI, native menus, and quit dialogs. Code editors (Monaco), integrated terminals (xterm), and AI conversation streams remain strictly untouched.
- **Offline & Reliable**: Zero network activity, zero data collection. Equipped with upfront JavaScript syntax validation and transactional rollback on unexpected errors.

---

## Download & Usage

### Download

Download the latest `antigravity-i18n.exe` from the [Releases](../../releases) page and run it directly.

### 1. Graphical Interface (Recommended)

Double-click `antigravity-i18n.exe`:
1. The app automatically detects the Antigravity installation folder (or click "Change Location" to specify manually).
2. Click on your preferred language (selecting **English** restores the official build).
3. Click "Apply"; the tool confirms process termination, updates the resources, and relaunches Antigravity.

### 2. Command Line Interface

The same binary can be invoked from the terminal:

```powershell
# Check current localization status and backups
antigravity-i18n.exe status

# Install a language pack (Simplified Chinese by default)
antigravity-i18n.exe apply --locale zh-CN

# Restore the official untranslated English version
antigravity-i18n.exe restore

# List all bundled language packs with coverage
antigravity-i18n.exe locales
```

Common options:
- `--app-dir <path>`: Manually specify the Antigravity installation directory (or set via `ANTIGRAVITY_APP_DIR`).
- `--locale <code>`: Target language code (e.g., `zh-CN`, `zh-Hant`, `ja`, `ko`, `de`, `es`, `fr`).
- `--no-restart`: Do not relaunch Antigravity after patching.
- `--force`: Skip the 30-second graceful shutdown wait and terminate processes immediately.
- `--yes`: Suppress the shutdown confirmation prompt.

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
- `app.asar.clean-backup`: The pristine official backup required for restoration. **Do not delete manually.**
- `app.asar.bak-*`: Historical timestamped snapshots left beside the archive. Once your current installation is verified stable, you may safely delete these files to free up disk space.

### 4. Why do some words or sentences remain in English?
Translation matches complete text rather than individual words in an unknown sentence. Fragments split by inline markup are also left intact to avoid partially translated sentences. Independent buttons, menu items, and accessibility labels can still be translated.

Voice mode preserves user speech, word-highlighted replies, and transcript text, while translating voice controls and work status. The Chinese packs retain the permission preset name `Turbo`, whose behavior concerns automatic execution and access permissions rather than speed alone.

---

## Building from Source

To compile from source:

**Prerequisites**:
- Go 1.26+
- [Wails v2](https://wails.io/) CLI

```powershell
# Clone the repository
git clone https://github.com/wadewu-ml/antigravity-i18n.git
cd antigravity-i18n

# Build both GUI and CLI binaries
.\build.ps1

# Run tests (including end-to-end asar verification)
$env:ANTIGRAVITY_ASAR = "$env:LOCALAPPDATA\Programs\antigravity\resources\app.asar.clean-backup"
go test ./...
```

DOM regression tests use an isolated headless Chrome or Edge instance to check voice content, sentence fragments, UI labels, and live updates. Set `ANTIGRAVITY_TEST_BROWSER` to choose another Chromium executable. These tests are skipped when no browser is available.

---

## Disclaimer

1. This is an unofficial, community-maintained project and is not affiliated with, endorsed by, or authorized by Google or Antigravity.
2. Google Antigravity and related trademarks belong to their respective owners.
3. This tool is provided for personal and educational use. Please ensure compliance with relevant software terms of service.

## License

[MIT License](./LICENSE)
