# antigravity-i18n


[Google Antigravity](https://antigravity.google/) デスクトップ クライアント用の UI ローカリゼーション ツール。


[简体中文](./README.zh-CN.md) | [English](./README.md) | [繁體中文](./README.zh-Hant.md) | **日本語** | [한국어](./README.ko.md) | [Deutsch](./README.de.md) | [Español](./README.es.md) | [Français](./README.fr.md) | [Português (Brasil)](./README.pt-BR.md) | [Русский](./README.ru.md)

---

**v4.0.0** は 9 つの言語パックを収録し、Antigravity App **2.33.0** で動作を確認しています。

---

## 主な機能

- **Windows x64 向けのポータブル実行ファイル**。グラフィカルインターフェイスは Microsoft Edge WebView2 を使用します。
- **GUI と CLI の両対応**: デフォルトで直感的なデスクトップ画面を起動。自動化やスクリプト向けに完全な CLI サブコマンド (`apply` / `restore` / `status` / `locales`) も内包しています。
- **バイト単位の正確な復元**: 初回変更前に公式オリジナルのバックアップ (`app.asar.clean-backup`) を自動作成。復元後は公式アーカイブと SHA-256 ハッシュが完全に一致します。
- **インターフェイスの翻訳**。シェル UI、ネイティブメニュー、終了ダイアログを対象とします。
- **オプションのセッション指標**。モデル選択の右側にリングでコンテキスト使用量を表示します。詳細には残りのコンテキスト、出力速度、キャッシュヒット率、最初のトークンまでの時間を表示します。Antigravity がローカルのセッションデータベースから値を読み取ります。
- **ローカルでのパッチ処理**。JavaScript の構文検証を行い、書き込みに失敗した場合はロールバックします。

---

## ダウンロードと使用方法

### ダウンロード

リポジトリの [Releases](../../releases) ページから最新の `antigravity-i18n.exe` をダウンロードして直接実行します。

### 1. グラフィカル インターフェイス (推奨)

`antigravity-i18n.exe` をダブルクリックします:
1. Antigravity のインストール先フォルダーを自動検出します (未検出時は「場所の変更」をクリック)。
2. 表示言語を選び、セッション指標を個別に設定します。公式の英語インターフェイスでも指標を表示できます。
3. 「適用」をクリックします。Antigravity が終了すると、リソースが更新され、アプリが再起動します。 「公式版に復元」で公式の英語アーカイブを復元し、指標を無効にできます。

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
- `apply --session-metrics` で指標を有効にし、`--session-metrics=false` で無効にします。`apply --locale en --session-metrics` で公式の英語インターフェイスに指標を表示します。
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
