# antigravity-i18n


[Google Antigravity](https://antigravity.google/) デスクトップ クライアント用の UI ローカリゼーション ツール。


[简体中文](./README.zh-CN.md) | [English](./README.md) | [繁體中文](./README.zh-Hant.md) | **日本語** | [한국어](./README.ko.md) | [Deutsch](./README.de.md) | [Español](./README.es.md) | [Français](./README.fr.md) | [Português (Brasil)](./README.pt-BR.md) | [Русский](./README.ru.md)

---

最新のグラフィカル インターフェイス、純粋な Go asar エンジン、および完全な言語パックを内蔵したスタンドアロン実行可能ファイル。Node.js や外部ランタイム不要でそのまま動作し、いつでも公式ビルドへバイト単位で正確に復元できます。

---

## 主な機能

- **単一ファイル・インストール不要**: 約 12 MB の独立した実行可能ファイル。Node.js、npm、Python 等の環境構築は一切不要です。
- **GUI と CLI の両対応**: デフォルトで直感的なデスクトップ画面を起動。自動化やスクリプト向けに完全な CLI サブコマンド (`apply` / `restore` / `status` / `locales`) も内包しています。
- **バイト単位の正確な復元**: 初回変更前に公式オリジナルのバックアップ (`app.asar.clean-backup`) を自動作成。復元後は公式アーカイブと SHA-256 ハッシュが完全に一致します。
- **非侵入型設計**: シェル UI、ネイティブ メニュー、終了ダイアログのみを変更。コード エディター (Monaco)、ターミナル (xterm)、AI 会話ストリームは完全にオリジナルのまま保護されます。
- **オフライン & 安全**: ネットワーク通信やデータ収集は一切行いません。JavaScript 構文検証とトランザクション ロールバックを備え、エラー時も安全に復元されます。

---

## ダウンロードと使用方法

### ダウンロード

リポジトリの [Releases](../../releases) ページから最新の `antigravity-i18n.exe` をダウンロードして直接実行します。

### 1. グラフィカル インターフェイス (推奨)

`antigravity-i18n.exe` をダブルクリックします:
1. Antigravity のインストール先フォルダーを自動検出します (未検出時は「場所の変更」をクリック)。
2. 言語一覧から希望の言語を選択します (**English** を選択すると公式原版に復元)。
3. 「適用」をクリックします。Antigravity が終了すると、リソースが更新され、アプリが再起動します。

適用や復元で Antigravity を閉じる必要がある場合、選択したインストール先のアプリに通常終了を要求します。30 秒経っても終了しない場合は、手動で終了して再試行するか、「強制終了して再試行」をクリックします。

### 2. コマンドライン モード

`antigravity-i18n.exe` のあるフォルダーで PowerShell を開きます。

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

一般的なオプション:
- `--app-dir <path>`: Antigravity のインストール ディレクトリを手動指定 (環境変数 `ANTIGRAVITY_APP_DIR` も利用可能)。
- `--locale <code>`: 対象言語コード (`zh-CN`、`zh-Hant`、`ja`、`ko`、`de`、`es`、`fr` 等)。
- `--no-restart`: 適用後に Antigravity を自動再起動しない。
- `--force` を指定すると、選択したインストール先の Antigravity を直ちに強制終了します。未保存の作業が失われることがあります。
- `--timeout <ms>` で操作を中止するまでの待ち時間をミリ秒単位で指定できます。既定値は `30000` です。
- `--yes`: 終了確認プロンプトに自動同意。

同じインストール先では、適用または復元を一度に一つだけ実行できます。別の操作が進行中の場合は、完了を待ってから再試行してください。`status` でインストール先とバックアップを確認できます。アーカイブを読み込めない場合や Antigravity のものと確認できない場合は、エラーが表示されます。

---

## 内蔵言語サポート

すべての言語パックはバイナリに直接埋め込まれています:

| コード | 言語 | カバレッジ | 説明 |
| :--- | :--- | :--- | :--- |
| `zh-CN` | Simplified Chinese (简体中文) | 100% | シェル UI、ネイティブ メニュー、およびダイアログの完全なカバレッジ |
| `zh-Hant` | Traditional Chinese (繁體中文) | 100% | 検証済みの技術用語による完全なカバレッジ |
| `ja` | Japanese (日本語) | 100% | 検証済みの技術用語による完全なカバレッジ |
| `ko` | Korean (한국어) | 100% | 検証済みの技術用語による完全なカバレッジ |
| `de` | German (Deutsch) | 100% | 検証済みの技術用語による完全なカバレッジ |
| `es` | Spanish (Español) | 100% | 検証済みの技術用語による完全なカバレッジ |
| `fr` | French (Français) | 100% | 検証済みの技術用語による完全なカバレッジ |
| `pt-BR` | Portuguese (Brasil) | 100% | 検証済みの技術用語による完全なカバレッジ |
| `ru` | Russian (Русский) | 100% | 検証済みの技術用語による完全なカバレッジ |
| `en` | English | - | 公式の未変更ビルド (復元するにはこれを選択します) |

---

## 仕組みとよくある質問

### 1. 公式アップデート後の対応
公式アップデートにより `app.asar` が上書きされます。更新完了後、本ツールを再度開いて「適用」をクリックするだけで再有効化されます (ツールが自動的に新バージョン基準でバックアップを更新します)。

### 2. 言語切り替えと復元
異なる言語間は事前の復元なしに直接切り替え可能です。いつでも「公式英語に戻す」を選択して初期状態へ復元できます。

### 3. バックアップ ファイルの管理
- `app.asar.clean-backup` は復元に使う元のアーカイブです。パッチが未適用の状態で適用または復元を行うと、インストール済みアーカイブとの SHA-256 を照合します。同じバージョンを再インストールした場合でも、ハッシュが異なればバックアップを更新します。**このファイルを保持してください。**
- `app.asar.bak-*` はタイムスタンプ付きのコピーです。元のバックアップがない場合や無効な場合に復旧に使えます。アプリの動作を確認し、`status` でバックアップが有効と確認できたら、古いコピーを削除して容量を確保できます。

---

## ソースからのビルド

**必要な環境**:
- Go 1.26.3+
- [Wails v2](https://wails.io/) CLI
- PowerShell 7

```powershell
# Clone the repository
git clone https://github.com/wadewu-ml/antigravity-i18n.git
cd antigravity-i18n

# Build both GUI and CLI binaries
.\build.ps1

# 公式アーカイブの復元テストを含める場合は次の行のコメントを外す
# $env:ANTIGRAVITY_ASAR = "$env:LOCALAPPDATA\Programs\antigravity\resources\app.asar.clean-backup"
go test ./...
```

---

## 免責事項

1. 本プロジェクトはコミュニティにより保守される非公式ツールであり、Google または Antigravity 公式とは一切関係ありません。
2. Google Antigravity および関連商標の著作権は各権利者に帰属します。
3. 本ツールは学習・個人利用を目的として提供されます。関連する利用規約を遵守してご使用ください。

## ライセンス

[MIT License](./LICENSE)
