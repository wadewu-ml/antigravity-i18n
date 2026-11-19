# antigravity-i18n

适用于 [Google Antigravity](https://antigravity.google/) 桌面客户端的多语言本地化工具。

**简体中文** | [English](./README.md) | [繁體中文](./README.zh-Hant.md) | [日本語](./README.ja.md) | [한국어](./README.ko.md) | [Deutsch](./README.de.md) | [Español](./README.es.md) | [Français](./README.fr.md) | [Português (Brasil)](./README.pt-BR.md) | [Русский](./README.ru.md)

---

单个可执行文件集成现代化图形界面、纯 Go 实现的 asar 读写引擎与完整语言包；无需安装 Node.js 或任何运行库，解压即用，并支持随时按字节原样还原为官方原版。

---

## 兼容性

当前源码已适配 Antigravity App 2.33.0。九套语言包补齐了浏览数据管理、快捷键方案、内置斜杠命令说明、SSH 主机操作、GitHub 连接错误、侧问操作、插件组合包及语音听写状态。应用、切换语言和还原已用官方 2.33.0 归档验证。参见 [Antigravity 官方更新记录](https://antigravity.google/docs/changelog#rel-hub-2.33.0)。

此前的设置、首次使用等界面适配见 [2.32.1 核对记录](./docs/translation-audit-2.32.1.md)。翻译引擎保留语音转写、通话摘要、AI 进度正文和侧问中的用户文字。翻译说明时也会保留链接、控件及主机名和搜索词等内容。简繁中文保留权限预设名称 Turbo。词库对齐率按内置语言包的词条计算，适配范围和限制见核对记录。

---

## 核心特性

- **单文件免安装**：约 12 MB 独立绿色程序，无需配置 Node.js、npm 或 Python 环境，开箱即用。
- **图形与命令行双支持**：默认启动直观易用的图形界面，同时内嵌完整命令行子命令（`apply` / `restore` / `status` / `locales`），兼顾日常使用与脚本自动化。
- **字节级原样还原**：首次修改前自动保存官方原始归档（`app.asar.clean-backup`）；执行还原后，`app.asar` 与官方发布的初始文件 SHA-256 哈希完全一致。
- **非侵入式设计**：仅本地化外壳 UI、原生系统菜单与退出确认框；代码编辑器（Monaco）、集成终端（xterm）与 AI 对话区域完全保持原生，不触碰任何用户数据。
- **离线与高可靠**：全流程零网络请求、零数据上传；内置脚本语法合法性门禁与全流程事务控制，遇到异常自动安全回滚。

---

## 下载与使用

### 下载运行

前往项目的 [Releases](../../releases) 页面，下载最新版本的 `antigravity-i18n.exe` 即可直接使用。

### 1. 图形界面（推荐）

双击运行 `antigravity-i18n.exe`：
1. 程序会自动检测 Antigravity 的安装目录（若未检测到，可点击“更改位置”手动选择）。
2. 在语言列表中点击选中目标语言（选中 **English** 即为还原官方原版）。
3. 点击“应用”并确认。Antigravity 退出后，工具会改写资源并重启客户端。

应用和还原需要关停客户端时，工具会请求选定安装的 Antigravity 正常退出。等待 30 秒后仍在运行，本次操作取消。需要立即关闭时，可在确认窗口勾选“强制退出”。

### 2. 命令行模式

在 `antigravity-i18n.exe` 所在目录打开 PowerShell。

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

常用参数：
- `--app-dir <path>`：手动指定 Antigravity 安装路径（亦可通过环境变量 `ANTIGRAVITY_APP_DIR` 设置）。
- `--locale <code>`：指定语言代码，如 `zh-CN`、`zh-Hant`、`ja`、`ko`、`de`、`es`、`fr` 等。
- `--no-restart`：完成改写后不自动重启客户端。
- 使用 `--force` 立即结束选定安装的 Antigravity 进程。未保存的工作可能丢失。
- 使用 `--timeout <ms>` 设置退出等待时长，单位为毫秒，默认 `30000`。
- `--yes`：静默确认进程关闭提示。

同一安装同一时间只能执行一次应用或还原。收到另一项操作正在运行的提示时，等它结束后再试。用 `status` 检查所选安装及备份。归档无法读取或应用身份无法验证时，命令会报错。

---

## 内置语言支持

所有语言包均已内嵌于二进制中，无需额外下载：

下表的覆盖度表示与内置基准词库的条目对齐程度，不代表客户端所有英文都应被翻译。语音正文、用户内容和无法完整匹配的句子片段会保留原文。

| 语言代码 | 语言名称 | 翻译覆盖度 | 说明 |
| :--- | :--- | :--- | :--- |
| `zh-CN` | 简体中文 | 100% | 完整覆盖界面外壳、原生菜单与退出弹窗 |
| `zh-Hant` | 繁體中文 | 100% | 完整覆盖，标准繁体计算机术语 |
| `ja` | 日本語 | 100% | 完整覆盖，专业 IDE 术语核对与对齐 |
| `ko` | 한국어 | 100% | 完整覆盖，专业 IDE 术语核对与对齐 |
| `de` | Deutsch | 100% | 完整覆盖，专业 IDE 术语核对与对齐 |
| `es` | Español | 100% | 完整覆盖，专业 IDE 术语核对与对齐 |
| `fr` | Français | 100% | 完整覆盖，专业 IDE 术语核对与对齐 |
| `pt-BR` | Português (Brasil) | 100% | 完整覆盖，专业 IDE 术语核对与对齐 |
| `ru` | Русский | 100% | 完整覆盖，专业 IDE 术语核对与对齐 |
| `en` | English | - | 官方原版英文（选中即执行还原） |

---

## 机制与常见问题

### 1. Antigravity 官方更新后失效？
Antigravity 自动更新会覆盖 `app.asar`。官方更新完成后，重新打开本工具点击一次“应用”即可重新生效；工具会自动基于新版文件刷新干净备份。

### 2. 语言切换与还原
不同语言之间可以直接点击切换，工具会自动替换现有补丁，无需先手动执行还原。任何时候点击“还原为官方英文”均可恢复至纯净状态。

### 3. 备份文件说明
- `app.asar.clean-backup` 保存用于还原的原始归档。对未打补丁的安装执行应用或还原时，工具会比对它与当前归档的 SHA-256。即使版本号相同，哈希不一致时也会刷新备份。**请保留此文件。**
- `app.asar.bak-*` 是带时间戳的归档快照。干净备份缺失或无效时，工具会从这些快照查找可用原档。确认客户端运行正常且 `status` 显示干净备份有效后，可清理旧快照释放空间。

### 4. 为什么部分词语或句子保持英文？
翻译按完整文本匹配，不对未知英文句子逐词替换。句子被内联标签拆分时，也会跳过片段，避免局部中文混入英文。独立按钮、菜单项和无障碍标签仍可正常翻译。

语音模式的用户发言、逐词高亮回复和转写正文保持原文，语音按钮与工作状态仍可翻译。简繁中文保留 `Turbo` 这一权限模式名称；它涉及自动执行和访问权限，译成“极速”容易产生误解。

---

## 源码构建

如需自行编译：

**依赖环境**：
- Go 1.26.3+
- [Wails v2](https://wails.io/) CLI
- PowerShell 7

```powershell
# 克隆仓库
git clone https://github.com/wadewu-ml/antigravity-i18n.git
cd antigravity-i18n

# 完整构建（同时产出 GUI 与 CLI）
.\build.ps1

# 取消下一行注释可一并验证官方归档的逐字节还原
# $env:ANTIGRAVITY_ASAR = "$env:LOCALAPPDATA\Programs\antigravity\resources\app.asar.clean-backup"
go test ./...
```

DOM 回归使用独立的无头 Chrome 或 Edge 实例。可用 `ANTIGRAVITY_TEST_BROWSER` 指定 Chromium 可执行文件。未找到浏览器时会跳过 DOM 用例。

---

## 免责声明

1. 本项目为开源社区维护的非官方本地化工具，与 Google 或 Antigravity 官方无关。
2. Antigravity 及相关商标版权归原权利人所有。
3. 本工具仅供学习交流及个人使用，请在遵循相关软件使用协议的前提下合理使用。

## 开源协议

[MIT License](./LICENSE)
