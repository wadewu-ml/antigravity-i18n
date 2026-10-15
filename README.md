# antigravity-i18n

给 [Antigravity](https://antigravity.google/) 桌面客户端换一套中文界面，想还原的时候按字节还原回官方原版。

一个 EXE 装下所有东西：界面、9 种语言包、读写 asar 的引擎。不装 Node、不装 npm、运行时不联网。

同一个 EXE 既开图形窗口，也认 `apply / restore / status / locales` 子命令。

## 为什么是 Go

最早这工具是 Node 写的，跑起来得带着一个 Node 运行时，打出来的包体积也大。这一版整个用 Go 重写，语言包和前端全部 `//go:embed` 进二进制，编译出来就是单个可执行文件，约 12 MB。

顺带把最费事的那一步也换掉了。原先是调 `@electron/asar` 的 `extractAll`：把整个 asar 解出来，改完再整个打回去，一次动几万个文件。Go 版直接改 asar 头部——读 header 拿到每个文件的 offset 和 size，按偏移把原来的字节整段拷过去，只替换需要改的三个文件。

## 用法

### 图形界面

双击 `antigravity-i18n.exe`。选语言，点按钮，剩下的事它自己做：先请 Antigravity 关闭（会先弹确认框），改写 app.asar，再把它启动回来。

语言列表里第一个 English 不是语言包，是官方原版。选中它点按钮就是还原。

### 命令行

```powershell
# 看当前状态，不需要关 Antigravity
antigravity-i18n.exe status

# 装简体中文（默认语言）
antigravity-i18n.exe apply --locale zh-CN

# 还原成官方英文
antigravity-i18n.exe restore

# 列出所有语言
antigravity-i18n.exe locales
```

| 选项 | 说明 |
| --- | --- |
| `--app-dir <path>` | 指定 Antigravity 安装目录，也可用环境变量 `ANTIGRAVITY_APP_DIR` |
| `--locale <code>` | 语言代码，`zh-CN`、`zh-Hant`、`ja`、`de` 等 |
| `--no-restart` | 改完不自动重启 Antigravity |
| `--force` | 跳过等待，直接结束进程 |
| `--timeout <ms>` | 等待正常退出的毫秒数，默认 30000 |
| `--yes` | 不再询问"是否关闭 Antigravity" |

找不到安装位置时，按 `ANTIGRAVITY_APP_DIR`、`%LOCALAPPDATA%\Programs\antigravity`、`Program Files`、`Program Files (x86)` 的顺序找。

## 改了什么

只替换三个文件：`dist/main.js`、`dist/menu.js`、`dist/preload.js`。其余 700 多个文件从原 asar 逐字节拷贝，偏移和 integrity 全部重算。

翻译范围是界面外壳、原生菜单、退出确认框。编辑器（Monaco）、终端、对话正文一律不碰——那些地方是你自己的内容，不该被词典改写。

`app.asar.unpacked/` 里的原生模块不参与 offset 计算，一个字节都不动。

写入时有个细节值得说：header 的 JSON 不重新序列化。Go 的 `encoding/json` 会按自己的规则排 key，和 Electron 写出来的对不上。这里是在原始 header 字节上做定位修改，key 顺序、缩进、转义都保持原样。

## 还原和备份

第一次 apply 之前，先存一份干净的 `app.asar.clean-backup`。`restore` 就是把它放回去，放回去之后 app.asar 的 SHA-256 和备份完全一致。

每次 apply/restore 还会在旁边留一个带时间戳的快照（`app.asar.bak-<时间>`，还原时叫 `app.asar.bak-before-restore-<时间>`）。这些会越积越多，装好确认没问题之后可以自己删——restore 真正需要的是 `app.asar.clean-backup` 那一个。

Antigravity 自己更新后会把 app.asar 覆盖掉，旧的 clean backup 就和新版对不上了。这种情况再跑一次 apply 就行：工具会拿当前这个未打补丁的 asar 重新生成干净备份。apply/restore 期间会比对哈希，发现 asar 中途被改动（比如自动更新插进来）就停下，不硬写。

备份是否可用由 `status` 判断，校验的是 package.json 的身份（name/version）、有没有被动过补丁、每个文件的 offset/size 有没有越界、integrity 对不对。不合格的备份不会被当成可还原的来源。

## 注意

1. 动手前先存盘。工具会请 Antigravity 正常退出并最多等 30 秒，超时强制结束——没保存的编辑器内容、会话、撤销历史会丢。CLI 的 `--yes` 和 GUI 的确认框都挡在这一步前面。
2. 语言之间可以直接切换，不需要先 restore。
3. 不做自动更新：没有版本检查，也没有下载。

## 开发

### 构建

版本号只有一处来源：仓库根目录的 `VERSION`。`build.ps1` 读它，用 `-ldflags` 注入 Go 二进制，同时写进 `wails.json` 的 `productVersion`（Wails 拿它填 PE 资源里的文件版本）。CLI 的 `--version`、窗口里的版本徽章、状态文件记的 `toolVersion` 因此始终一致。

```powershell
.\build.ps1              # GUI + CLI
.\build.ps1 -Target cli  # 只要 CLI
```

不经脚本直接 `go build` 不会注入版本，运行时报告 `dev`。

### 测试

e2e 用本机真实的 asar 跑完整 apply → restore 闭环，所以要指一份干净的 asar：

```powershell
$env:ANTIGRAVITY_ASAR = "$env:LOCALAPPDATA\Programs\antigravity\resources\app.asar.clean-backup"
go test ./...
```

### 代码结构

| 路径 | 作用 |
| --- | --- |
| `internal/core/asar.go` | asar 读写：header 解析、字节拷贝、offset 与 integrity 重算 |
| `internal/core/patch.go` | 三个 JS 文件的注入逻辑（哨兵块、锚点、去重） |
| `internal/core/locale.go` | 语言包加载与结构校验 |
| `internal/core/installer.go` | 备份链、clean backup 校验、状态判定 |
| `internal/core/apply.go` | apply / restore 主流程 |
| `internal/core/txn.go` | 事务与回滚 |
| `frontend/dist/` | 手写静态界面（HTML/CSS/JS，没有构建步骤） |
| `app.go` | Wails 绑定，GUI 和 CLI 共用同一套 core |

注入的 JS 会先用 goja 解析器过一遍语法，不合格就整个操作失败，不会留下打不开的界面。

## 版本

Go 线从 `3.0.0` 起。`v1.0.0` 到 `v2.1.0` 这几个 tag 属于归档的 Node 实现，版本号不重用，所以不冲突。语言包词条同样来自那份实现；要参考它，可以直接检出这些 tag。

## 许可证

MIT（沿用原项目）。社区维护的非官方本地化工具，与 Antigravity 官方无关。
