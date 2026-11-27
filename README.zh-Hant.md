<div align="center">
  <img src="./frontend/dist/appicon.png" width="68" height="68" alt="Regravity 圖示">
  <h1>Regravity</h1>
  <p><em>反重力起飛，Regravity 讓它落地——回歸你的語言與掌控。</em></p>
  <p>可攜式工具 Regravity 為 Windows 版 <strong>Google Antigravity</strong> 提供 15 個翻譯套件，並提供會話指標、模型額度與思考強度滑桿。Regravity 設定視窗支援 16 種語言。</p>
  <p>
    <a href="https://github.com/wadewu-ml/regravity/releases/latest/download/regravity.exe"><img src="https://img.shields.io/badge/Download-Latest_(Windows_x64)-0F9D58?style=flat&logo=windows&logoColor=white" alt="下載最新版 Windows x64 Regravity"></a>
    <a href="https://github.com/wadewu-ml/regravity/releases"><img src="https://img.shields.io/badge/Releases-GitHub-18181B?style=flat&logo=github" alt="GitHub 版本記錄"></a>
    <img src="https://img.shields.io/badge/Antigravity-2.35.2_tested-18181B?style=flat" alt="已驗證適用 Antigravity App 2.35.2">
    <a href="./LICENSE"><img src="https://img.shields.io/badge/License-MIT-18181B?style=flat" alt="MIT 授權條款"></a>
  </p>
  <p>
    <sub>
      <a href="./README.zh-CN.md">简体中文</a> · <a href="./README.md">English</a> · <strong>繁體中文</strong> · <a href="./README.ja.md">日本語</a> · <a href="./README.ko.md">한국어</a> · <a href="./README.de.md">Deutsch</a> · <a href="./README.es.md">Español</a> · <a href="./README.fr.md">Français</a> · <a href="./README.pt-BR.md">Português (Brasil)</a> · <a href="./README.ru.md">Русский</a> · <a href="./README.tr.md">Türkçe</a> · <a href="./README.vi.md">Tiếng Việt</a> · <a href="./README.it.md">Italiano</a> · <a href="./README.ar.md">العربية</a> · <a href="./README.id.md">Bahasa Indonesia</a> · <a href="./README.pl.md">Polski</a>
    </sub>
  </p>
</div>

<br>

<p align="center">
  <a href="./docs/images/language-settings.webp">
    <img src="./docs/images/language-settings.webp" width="580" alt="Regravity 設定視窗，顯示 Antigravity v2.35.2、已選取簡體中文，並同時開啟會話指標與思考強度滑桿">
  </a>
  <br>
  <sub>Regravity 設定視窗，圖中用戶端為 Antigravity 2.35.2。</sub>
</p>

## 核心能力

- **15 個翻譯套件，不碰對話內容** — 翻譯設定、對話區控制項、模型選單、原生選單與結束提示，不改動使用者輸入與聊天紀錄。
- **會話指標與模型額度** *（可選）* — 在模型選擇器右側新增圓環，懸停即看各模型群組剩餘額度、本地重設時間、已用上下文、平均輸出速度、快取命中率與首個 token 延遲。
- **思考強度滑桿** *（可選）* — 在模型選單中為支援調節的模型嵌入膠囊滑桿，拖曳即切檔位。
- **獨立組合，逐位元組還原** — 兩項增強開關可分別開啟，也相容官方英文介面；套用後無需常駐背景，保留與目前 Antigravity 版本相符的乾淨備份即可逐位元組還原。

## 快速上手

> [!TIP]
> 括號內是上方截圖中的簡體中文按鈕名稱。

在右上角可以選擇設定視窗的語言。

### 套用設定

1. 按兩下執行 `regravity.exe`。若未自動找到正確安裝位置，點選 **更改位置**（`更改位置`）並選擇含有 `Antigravity.exe` 的資料夾。
2. 在 **介面語言**（`界面语言`）中選擇目標語言（若僅想保留英文介面並開啟可選功能，選 **English** 即可）。
3. 依需求開啟語言清單下方的 **會話指標**（`会话指标`）與 **思考強度滑桿**（`思考强度滑块`）。
4. 儲存工作並結束 Antigravity，點選右下角 **套用變更**（`应用更改`）。Antigravity 會以新設定重新開啟（也可點選 **開啟 Antigravity** / `打开 Antigravity` 啟動），之後即可關閉本工具。

> [!WARNING]
> 若提示 Antigravity 仍在執行，請先儲存工作後正常結束。使用 **強制結束並重試**（`强制退出并重试`）會中斷進行中的工作階段，尚未儲存的變更可能會遺失。

### 重新套用與還原

- **重新套用** — 更新工具後，點選 **重新套用**（`重新应用`）就能使用新譯文。
- **還原官方版本** — 先結束 Antigravity，開啟右上角選單（`···`），點選 **還原官方版本**（`还原官方版本`）並確認，介面將恢復官方英文並移除指標圓環與滑桿。

> [!IMPORTANT]
> 首次套用時，工具會在 Antigravity 安裝目錄自動備份原檔案至 `resources/app.asar.clean-backup`。如需日後還原，請保留該檔案。

## 可選功能

### 會話指標與模型額度

- **送出第一則訊息前即可查額度** — 游標移至模型選擇器右側圓環上，即可查看各模型群組在不同週期（如每週、5 小時）的剩餘額度；停在百分比上可查看本地重設時間。
- **產生模型呼叫後顯示即時指標** — 對話記錄到模型呼叫後，圓環反映估算上下文比例，面板上方自動填入 **已用上下文**（已用 token / 總容量與比例）、**平均輸出速度**（含模型記錄的思考 token）、**快取命中率** 與 **首個 token 延遲**。點一下圓環可固定或收起。
- **多面板獨立統計** — 分割畫面、側問和子代理面板各自讀取對應會話最近一次有資料的本機紀錄，無法取得的項目顯示為 `—`。

<p align="center">
  <a href="./docs/images/session-quota-before-first-call.webp">
    <img src="./docs/images/session-quota-before-first-call.webp" width="420" alt="模型選擇器右側圓環上方展開的會話指標與模型剩餘額度面板">
  </a>
  <br>
  <sub>新對話即可直接查看各模型群組剩餘額度；產生首次呼叫後上方指標自動填入。</sub>
</p>

### 思考強度滑桿

開啟 **思考強度滑桿**（`思考强度滑块`）並套用變更後，開啟 Antigravity 模型選單，支援直接選擇思考強度的模型會在用量選項上方顯示滑桿——拖曳檔位預覽，放開即生效。

<p align="center">
  <a href="./docs/images/thinking-effort-high.webp">
    <img src="./docs/images/thinking-effort-high.webp" height="340" alt="模型選單中的思考強度滑桿，目前為高">
  </a>
  &nbsp;&nbsp;
  <a href="./docs/images/thinking-effort-max.webp">
    <img src="./docs/images/thinking-effort-max.webp" height="340" alt="模型選單中的思考強度滑桿，目前為極高">
  </a>
  <br>
  <sub>常規檔位（高）與高亮最高檔位（極高），直接在模型選單中拖曳切換。</sub>
</p>

## 更新與命令列

### 更新說明

- **Antigravity 更新後** — 若翻譯或增強元件消失，請先對照頂端已驗證的用戶端版本，再重新開啟工具套用設定即可。
- **工具自動檢查更新** — 啟動時依上次檢查結果決定是否查詢 GitHub，成功後至少間隔 24 小時再查。右上角選單（`···`）出現藍點表示有新版本，點開可查看更新說明並就地下載重啟。可在 **關於**（`关于`）關閉自動檢查，或隨時手動點擊 **檢查更新**（`检查更新`）。工具更新後請點選 **重新套用**（`重新应用`）套用新譯文。

如果你使用 v4.2.3 或更早版本，請手動下載並執行 [`regravity.exe`](https://github.com/wadewu-ml/regravity/releases/latest/download/regravity.exe)，原有設定會保留。

### 命令列用法

在 `regravity.exe` 所在目錄開啟 PowerShell：

| 操作 | 指令 |
| --- | --- |
| 列出語言代碼 | `.\regravity.exe locales` |
| 套用繁體中文 | `.\regravity.exe apply --locale zh-Hant` |
| 官方英文 + 會話指標 | `.\regravity.exe apply --locale en --session-metrics` |
| 官方英文 + 思考強度滑桿 | `.\regravity.exe apply --locale en --effort-slider` |
| 查看目前設定 | `.\regravity.exe status` |
| 還原官方版本 | `.\regravity.exe restore` |
| 查看所有選項 | `.\regravity.exe --help` |

### 從原始碼建置

需要 Windows、PowerShell 7、Go 1.26.3+ 及 Wails v2 CLI：

```powershell
.\build.ps1 -Target gui
```

產物輸出至 `build/bin/regravity.exe`。

## 常見問題

- **啟動時提示缺少 WebView2** — 請安裝 [Microsoft Edge WebView2 執行階段](https://developer.microsoft.com/en-us/microsoft-edge/webview2/#download) 後重新執行 `regravity.exe`。

---

## 免責聲明

1. 本專案為開源社區維護的非官方本地化工具，與 Google 或 Antigravity 官方無關。
2. Antigravity 及相關商標版權歸原權利人所有。
3. 本工具僅供學習交流及個人使用，請在遵循相關軟體使用協議的前提下合理使用。

[MIT 授權條款](./LICENSE)
