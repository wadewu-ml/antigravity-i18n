# antigravity-i18n

適用于 [Google Antigravity](https://antigravity.google/) 桌面客戶端的多語言本地化工具。

[简体中文](./README.zh-CN.md) | [English](./README.md) | [繁體中文](./README.zh-Hant.md) | [日本語](./README.ja.md) | [한국어](./README.ko.md) | [Deutsch](./README.de.md) | [Español](./README.es.md) | [Français](./README.fr.md) | [Português (Brasil)](./README.pt-BR.md) | [Русский](./README.ru.md)

---

**v4.0.0** 已適配 Antigravity App **2.33.0**，內含九套語言包。

---

## 核心特性

- **Windows x64 單檔案**可直接執行。圖形介面使用 Microsoft Edge WebView2。
- **圖形與命令行雙支持**：預設啟動直觀易用的圖形界面，同時內嵌完整命令行子命令（`apply` / `restore` / `status` / `locales`），兼顧日常使用與腳本自動化。
- **字節級原樣還原**：首次修改前自動保存官方原始歸檔（`app.asar.clean-backup`）；執行還原后，`app.asar` 與官方發布的初始檔案 SHA-256 哈希完全一致。
- **介面本地化**涵蓋應用程式外殼、原生選單及結束對話框。
- **可選會話指標**在模型選擇器右側顯示上下文用量。懸停可查看剩餘上下文、輸出速度、快取命中率與首個 token 的延遲，點選可固定詳細資料。Antigravity 以唯讀方式查詢本機會話資料庫。
- **本機修補**包含 JavaScript 語法檢查，寫入失敗時會執行回復。

---

## 下載與使用

### 下載運行

前往專案的 [Releases](../../releases) 頁面，下載最新版本的 `antigravity-i18n.exe` 即可直接使用。

### 1. 圖形界面（推薦）

雙擊運行 `antigravity-i18n.exe`：
1. 程序會自動檢測 Antigravity 的安裝目錄（若未檢測到，可點擊“更改位置”手動選擇）。
2. 選擇介面語言，再單獨設定「會話指標」開關。官方英文介面也可顯示指標。
3. 按一下「應用」。Antigravity 結束後，工具會更新資源並重新啟動用戶端。 按一下「還原官方版本」可還原官方英文歸檔並關閉指標。

套用或還原需要關閉用戶端時，工具會要求所選安裝的 Antigravity 正常結束。等待 30 秒後若仍在執行，可以手動結束後再試，或按一下「強制結束並重試」。

### 2. 命令行模式

在 `antigravity-i18n.exe` 所在資料夾開啟 PowerShell。



```powershell
# 查看当前本地化状态与备份信息
.\antigravity-i18n.exe status

# 安装指定语言包（默认简体中文）
.\antigravity-i18n.exe apply --locale zh-CN

# 随时还原为官方英文原版
.\antigravity-i18n.exe restore

# 查看所有内置语言包及其覆盖率
.\antigravity-i18n.exe locales
```



常用參數：
- `--app-dir <path>`：手動指定 Antigravity 安裝路徑（亦可通過環境變量 `ANTIGRAVITY_APP_DIR` 設定）。
- `--locale <code>`：指定語言代碼，如 `zh-CN`、`zh-Hant`、`ja`、`ko`、`de`、`es`、`fr` 等。
- `--no-restart`：完成改寫后不自動重啟客戶端。
- `apply --session-metrics` 開啟指標，`--session-metrics=false` 關閉。`apply --locale en --session-metrics` 可在官方英文介面開啟指標。
- 使用 `--force` 可立即結束所選安裝的 Antigravity 程序。未儲存的工作可能遺失。
- 使用 `--timeout <ms>` 設定等待結束的時間，單位為毫秒，預設為 `30000`。
- `--yes`：靜默確認進程關閉提示。

同一安裝同時只能執行一次套用或還原。若出現另一項操作正在執行的提示，請等候操作完成後再試。使用 `status` 檢查所選安裝與備份。歸檔無法讀取或無法驗證為 Antigravity 時，命令會回報錯誤。

---

## 內置語言支持

所有語言包均已內嵌于二進制中，無需額外下載：

| 語言代碼 | 語言名稱 | 翻譯覆蓋度 | 說明 |
| :--- | :--- | :--- | :--- |
| `zh-CN` | 簡體中文 | 100% | 完整覆蓋介面外殼、原生功能表與結束彈窗 |
| `zh-Hant` | 繁體中文 | 100% | 完整覆蓋，標準繁體電腦術語 |
| `ja` | 日本語 | 100% | 完整覆蓋，專業 IDE 術語核對與對齊 |
| `ko` | 한국어 | 100% | 完整覆蓋，專業 IDE 術語核对与对齐 |
| `de` | Deutsch | 100% | 完整覆蓋，專業 IDE 術語核对与对齐 |
| `es` | Español | 100% | 完整覆蓋，專業 IDE 術語核对与对齐 |
| `fr` | Français | 100% | 完整覆蓋，專業 IDE 術語核对与对齐 |
| `pt-BR` | Português (Brasil) | 100% | 完整覆蓋，專業 IDE 術語核对与对齐 |
| `ru` | Русский | 100% | 完整覆蓋，專業 IDE 術語核对与对齐 |
| `en` | English | - | 官方原版英文（選中即執行還原） |

---

## 機制與常見問題

### 1. Antigravity 官方更新后失效？
Antigravity 自動更新會覆蓋 `app.asar`。官方更新完成后，重新打開本工具點擊一次“應用”即可重新生效；工具會自動基于新版檔案刷新干凈備份。

### 2. 語言切換與還原
不同語言之間可以直接點擊切換，工具會自動替換現有補丁，無需先手動執行還原。任何時候點擊“還原為官方英文”均可恢復至純凈狀態。

### 3. 備份檔案說明
- `app.asar.clean-backup` 保存還原所需的原始歸檔。對未套用補丁的安裝執行套用或還原時，工具會比對備份與目前歸檔的 SHA-256 雜湊。即使版本相同，雜湊不同時也會更新備份。**請保留此檔案。**
- `app.asar.bak-*` 是有時間戳記的歸檔快照。乾淨備份遺失或無效時，工具會從這些快照尋找可用的原始歸檔。確認用戶端運作正常且 `status` 顯示備份有效後，可清除舊快照以釋放空間。

---

## 源碼構建

如需自行編譯：

**依賴環境**：
- Go 1.26.3+
- [Wails v2](https://wails.io/) CLI
- PowerShell 7



```powershell
# 克隆仓库
git clone https://github.com/wadewu-ml/antigravity-i18n.git
cd antigravity-i18n

# 完整构建（同时产出 GUI 与 CLI）
.\build.ps1

# 取消下一行註解可一併驗證官方歸檔的逐位元組還原
# $env:ANTIGRAVITY_ASAR = "$env:LOCALAPPDATA\Programs\antigravity\resources\app.asar.clean-backup"
go test ./...
```



---

## 免責聲明

1. 本專案為開源社區維護的非官方本地化工具，與 Google 或 Antigravity 官方無關。
2. Antigravity 及相關商標版權歸原權利人所有。
3. 本工具僅供學習交流及個人使用，請在遵循相關軟體使用協議的前提下合理使用。

## 開源協議

[MIT License](./LICENSE)
