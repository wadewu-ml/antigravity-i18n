<div align="center">
  <img src="./frontend/dist/appicon.png" width="68" height="68" alt="Regravity アイコン">
  <h1>Regravity</h1>
  <p><em>Antigravity が飛び立つ先で、Regravity があなたの言語とコントロールを取り戻します。</em></p>
  <p>Windows 版 <strong>Google Antigravity</strong> 向けポータブル拡張ツール — 9 言語のローカライズ、セッション指標・利用枠パネル、推論強度スライダー。</p>
  <p>
    <a href="https://github.com/wadewu-ml/regravity/releases/latest/download/regravity.exe"><img src="https://img.shields.io/badge/Download-Latest_(Windows_x64)-0F9D58?style=flat&logo=windows&logoColor=white" alt="最新の Windows x64 版 Regravity をダウンロード"></a>
    <a href="https://github.com/wadewu-ml/regravity/releases"><img src="https://img.shields.io/badge/Releases-GitHub-18181B?style=flat&logo=github" alt="GitHub リリース一覧"></a>
    <img src="https://img.shields.io/badge/Antigravity-2.35.1_tested-18181B?style=flat" alt="Antigravity App 2.35.1 動作確認済み">
    <a href="./LICENSE"><img src="https://img.shields.io/badge/License-MIT-18181B?style=flat" alt="MIT ライセンス"></a>
  </p>
  <p>
    <sub>
      <a href="./README.zh-CN.md">简体中文</a> · <a href="./README.md">English</a> · <a href="./README.zh-Hant.md">繁體中文</a> · <strong>日本語</strong> · <a href="./README.ko.md">한국어</a> · <a href="./README.de.md">Deutsch</a> · <a href="./README.es.md">Español</a> · <a href="./README.fr.md">Français</a> · <a href="./README.pt-BR.md">Português (Brasil)</a> · <a href="./README.ru.md">Русский</a> · <a href="./README.tr.md">Türkçe</a> · <a href="./README.vi.md">Tiếng Việt</a>
    </sub>
  </p>
</div>

<br>

<p align="center">
  <a href="./docs/images/language-settings.webp">
    <img src="./docs/images/language-settings.webp" width="580" alt="Antigravity v2.35.2、簡体字中国語、セッション指標および推論強度スライダーが有効な設定画面">
  </a>
  <br>
  <sub>Regravity の設定画面。使用しているクライアントは Antigravity 2.35.2 です。</sub>
</p>

## 主な機能

- **9 言語収録・会話内容には干渉なし** — 設定、チャット操作、モデルメニュー、ネイティブメニュー、終了確認を `zh-CN` / `zh-Hant` / `ja` / `ko` / `de` / `es` / `fr` / `pt-BR` / `ru` に翻訳。入力内容や会話履歴は書き換えません。
- **セッション指標と残り利用枠** *（任意）* — モデル選択ボタン横にリングを追加し、残り利用枠、現地リセット時刻、コンテキスト使用量、平均出力速度、キャッシュヒット率、最初のトークンまでの時間を確認できます。
- **推論強度スライダー** *（任意）* — 推論強度を選べるモデルの選択メニュー内にカプセル型スライダーを組み込みます。
- **個別切替・バイト単位で復元** — 2 つの拡張スイッチは個別に切り替えでき、標準の英語表示とも併用可能。インストール済みの Antigravity と同じバージョンの未変更バックアップからいつでも復元できます。

## クイックスタート

> [!TIP]
> 括弧内はスクリーンショットに表示される中国語のラベルです。

開発版では右上から設定画面の言語を選べます。Antigravity の表示言語にトルコ語とベトナム語を追加しました。

### 設定を適用する

1. `regravity.exe` を開きます。インストール先が正しくない場合は **場所を変更**（`更改位置`）をクリックし、`Antigravity.exe` のあるフォルダーを選びます。
2. **表示言語**（`界面语言`）で言語を選びます（標準英語のまま拡張機能のみ使う場合は **English** を選択）。
3. 必要に応じて言語一覧の下にある **セッション指標**（`会话指标`）や **推論強度スライダー**（`思考强度滑块`）をオンにします。
4. 作業を保存して Antigravity を終了し、右下の **変更を適用**（`应用更改`）をクリックします。新しい設定で再起動したら（**Antigravity を開く** / `打开 Antigravity` からも起動可）、ツールは閉じてかまいません。

> [!WARNING]
> Antigravity が起動中という案内が出た場合は、作業を保存して通常どおり終了してください。**強制終了して再試行**（`强制退出并重试`）を選ぶと進行中のセッションが中断され、未保存の変更が失われることがあります。

### 再適用と復元

- **再適用** — ツールの更新後に **再適用**（`重新应用`）をクリックすると、新しい翻訳を使えます。
- **公式版に戻す** — Antigravity を終了し、右上メニュー（`···`）から **公式版に戻す**（`还原官方版本`）をクリックして確定すると、標準の英語表示に戻り指標リングとスライダーも削除されます。

> [!IMPORTANT]
> 初回適用時に Antigravity インストール先へ元のファイルが `resources/app.asar.clean-backup` としてバックアップされます。復元用にこのファイルを残しておいてください。

## 拡張機能

### セッション指標とモデルの残り利用枠

- **最初のメッセージ前から利用枠を確認** — モデル選択ボタン横のリングにポインターを合わせると、モデルグループごとの各期間（週間・5 時間など）の残り利用枠を表示。割合に合わせると現地リセット日時が確認できます。
- **モデル呼び出し後のリアルタイム指標** — 呼び出しが記録されるとリングが推定コンテキスト使用率を示し、上部に **コンテキスト使用量**（使用トークン / 総容量・割合）、**平均出力速度**（思考トークンを含む）、**キャッシュヒット率**、**最初のトークンまでの時間** が表示されます。クリックで固定・閉じられます。
- **パネルごとに正確に集計** — 分割画面、サイドチャット、サブエージェントはそれぞれのセッションのローカル記録から直近のモデル呼び出しを表示。取得できない項目は `—` になります。

<p align="center">
  <a href="./docs/images/session-quota-before-first-call.webp">
    <img src="./docs/images/session-quota-before-first-call.webp" width="420" alt="リングの上に開いたセッション指標とモデルの残り利用枠パネル">
  </a>
  <br>
  <sub>最初の送信前から利用枠を確認可能。モデル呼び出し後に上部の各指標が表示されます。</sub>
</p>

### 推論強度スライダー

**推論強度スライダー**（`思考强度滑块`）をオンにして適用すると、推論強度を選べるモデルのメニュー内（使用量項目の上）にスライダーが表示されます。ドラッグで段階を確認し、離すと適用されます。

<p align="center">
  <a href="./docs/images/thinking-effort-high.webp">
    <img src="./docs/images/thinking-effort-high.webp" height="340" alt="推論強度が「高」のモデル選択メニュー">
  </a>
  &nbsp;&nbsp;
  <a href="./docs/images/thinking-effort-max.webp">
    <img src="./docs/images/thinking-effort-max.webp" height="340" alt="推論強度が最大のモデル選択メニュー">
  </a>
  <br>
  <sub>通常レベル（高）とハイライトされた最大レベル（極高）をモデルメニュー内で直接切替。</sub>
</p>

## アップデートと CLI

### アップデート

- **Antigravity の更新後** — 翻訳や拡張機能が消えた場合は、このページ上部の動作確認済みバージョンを確認のうえ、再度ツールから適用してください。
- **ツールの自動更新** — 起動時に前回の結果をもとに GitHub への確認を判断し、成功後は最低 24 時間間隔を空けます。右上メニュー（`···`）の青い点から変更内容の確認とインプレース更新が可能。自動確認は **このツールについて**（`关于`）で切替でき、手動確認（`检查更新`）も利用できます。更新後は **再適用**（`重新应用`）で新しい翻訳を反映してください。

v4.2.3 以前をお使いの場合は、[`regravity.exe`](https://github.com/wadewu-ml/regravity/releases/latest/download/regravity.exe) を手動でダウンロードして実行してください。設定は引き継がれます。

### PowerShell CLI

`regravity.exe` のあるフォルダーで PowerShell を開いて実行します：

| 操作 | コマンド |
| --- | --- |
| 言語コードを一覧表示 | `.\regravity.exe locales` |
| 日本語を適用 | `.\regravity.exe apply --locale ja` |
| 標準英語 + セッション指標 | `.\regravity.exe apply --locale en --session-metrics` |
| 標準英語 + 推論強度スライダー | `.\regravity.exe apply --locale en --effort-slider` |
| 現在の設定を確認 | `.\regravity.exe status` |
| 公式版に復元 | `.\regravity.exe restore` |
| 全オプションを表示 | `.\regravity.exe --help` |

### ソースからビルドする

Windows、PowerShell 7、Go 1.26.3 以降、Wails v2 CLI が必要です：

```powershell
.\build.ps1 -Target gui
```

成果物は `build/bin/regravity.exe` に出力されます。

## トラブルシューティング

- **起動時に WebView2 が必要と表示される** — [Microsoft Edge WebView2 ランタイム](https://developer.microsoft.com/en-us/microsoft-edge/webview2/#download) をインストールしてから `regravity.exe` を再起動してください。

---

## 免責事項

1. 本プロジェクトはコミュニティにより保守される非公式ツールであり、Google または Antigravity 公式とは一切関係ありません。
2. Google Antigravity および関連商標の著作権は各権利者に帰属します。
3. 本ツールは学習・個人利用を目的として提供されます。関連する利用規約を遵守してご使用ください。

[MIT ライセンス](./LICENSE)
