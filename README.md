# antigravity-i18n（Go 重写版）

给 [Antigravity](https://antigravity.google/) 桌面应用换一套中文界面，随后可以按字节原样还原成官方版本。

- **单文件**：一个 12 MB 的 EXE，包含全部界面、全部 9 种语言和 asar 读写引擎。
- **不需要 Node**：不依赖 Node/npm，不依赖运行时安装，也不需要在构建时装前端依赖。
- **界面为主**：双击打开图形窗口；同一个 EXE 里也内嵌了 `apply / restore / status / locales` 子命令。
- **按字节还原**：第一次应用前会留下一份干净的 `app.asar.clean-backup`；`restore` 之后 asar 的 SHA-256 与备份完全一致。
- **只动三个文件**：`dist/main.js`、`dist/menu.js`、`dist/preload.js`，其余 700 多个文件从原 asar 逐字节拷贝，偏移与 integrity 全部重算。
- **离线**：没有任何网络请求，不读取 token、会话或凭据。

界面语言包取自原项目 `D:\WS\Selene-WS\projects\antigravity-zh`，中文覆盖 100%，其余语言沿用其词条集合。原项目保持只读参考。

---

## 使用

### 图形界面

直接运行 `antigravity-i18n.exe`。选择语言后点“应用中文”，工具会自动关闭 Antigravity、改写 `app.asar`、再把它启动回来。

### 命令行

```powershell
# 查看当前状态（不需要关闭 Antigravity）
antigravity-i18n.exe status

# 应用简体中文（默认就是 zh-CN）
antigravity-i18n.exe apply --locale zh-CN

# 还原为官方英文
antigravity-i18n.exe restore

# 列出所有语言
antigravity-i18n.exe locales
```

常用选项：

| 选项 | 说明 |
| --- | --- |
| `--app-dir <path>` | 指定 Antigravity 安装目录；也可用环境变量 `ANTIGRAVITY_APP_DIR` |
| `--locale <code>` | 语言代码，如 `zh-CN`、`ja`、`de` |
| `--no-restart` | 应用完不自动重启 Antigravity |
| `--force` | 跳过优雅关闭，直接结束进程 |
| `--timeout <ms>` | 等待优雅关闭的毫秒数，默认 30000 |
| `--yes` | 不再询问“是否关闭 Antigravity” |

---

## 与 Node 版的区别

原实现用 Node 跑 `@electron/asar` 的 `extractAll` + 重新打包：把整个 asar 解出来再压回去，一次读写几万个文件，还要求目标机器装 Node。

Go 版改为**直接改 asar 头部再拷贝字节**：

1. 读 header，取得每个文件的 `offset` / `size` / `integrity`。
2. 按 `offset` 把原文件的字节整段拷到新 asar，只替换三个 JS 文件的内容。
3. 重算被替换文件的 `size`、后续文件的 `offset`、以及它们的 SHA-256 integrity。
4. header JSON 不做重新序列化，只在原文上做定位修改，因此 key 顺序、缩进、转义都与 Electron 自己写的一致。

`app.asar.unpacked/` 里的原生模块（本机 347 个文件、13 MB）完全不受影响，也不进入 offset 计算。

---

## 安全边界

- 应用前会校验 `app.asar` 的 `package.json` 确实是 `antigravity`，避免把别的 app.asar 写坏。
- 应用前会关闭 Antigravity；替换前再检查一次进程和 asar 哈希，防止自动更新中途插入。
- 全流程带事务与回滚：任何一步失败都会把 `app.asar`、clean backup、状态文件恢复原状。
- 注入的 JS 会用 goja 解析器做语法检查，不合格就整个操作失败，不会留下打不开的界面。
- 只翻译界面外壳、原生菜单和退出确认框；编辑器、终端、对话区不动。

---

## 开发

```powershell
# 跑测试（需要本机真实 asar 做闭环验证）
$env:ANTIGRAVITY_ASAR = "$env:LOCALAPPDATA\Programs\antigravity\resources\app.asar.clean-backup"
go test ./...

# 构建（版本号取自 VERSION，同时产出 GUI 与 CLI）
.\build.ps1

# 只构建其中一个
.\build.ps1 -Target cli
.\build.ps1 -Target gui
```

版本号只有一处来源：仓库根目录的 `VERSION`。`build.ps1` 把它注入 Go 二进制
（`-ldflags -X antigravity-i18n/internal/core.Version=...`），并同步写入
`wails.json` 的 `productVersion`（Wails 用它填 PE 资源里的文件版本）。因此
CLI 的 `--version`、窗口标题栏的版本徽章、以及状态文件里记录的 `toolVersion`
始终是同一个数字。直接 `go build` 而不经脚本时不注入版本，运行时报告为
`dev`，不会冒充一个正式版本号。

GUI 与 CLI 是同一个 EXE 的两种形态：双击 = 图形界面，带子命令 = 命令行
（`apply` / `restore` / `status` / `locales` / `gui`）。

代码结构：

| 路径 | 作用 |
| --- | --- |
| `internal/core/asar.go` | asar 读写：header 解析、字节拷贝、offset 与 integrity 重算 |
| `internal/core/patch.go` | 三个 JS 文件的注入语义（哨兵块、锚点、去重） |
| `internal/core/locale.go` | 语言包加载与结构校验 |
| `internal/core/installer.go` | 备份链、clean backup 校验、状态判定 |
| `internal/core/apply.go` | apply / restore 主流程 |
| `internal/core/txn.go` | 事务与回滚 |
| `frontend/dist/` | 手写静态窗口（HTML/CSS/JS，无构建步骤） |
| `app.go` | Wails 绑定，GUI 与 CLI 共用同一套 core API |

## License

MIT（沿用原项目）。
