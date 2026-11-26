<div align="center">
  <img src="./frontend/dist/appicon.png" width="64" height="64" alt="Regravity 图标">
  <h1>Regravity</h1>
  <p>为 Windows 版 Google Antigravity 提供九种界面语言包、会话指标与模型额度面板，以及思考强度滑块。</p>
  <p>
    <a href="https://github.com/wadewu-ml/regravity/releases/download/v4.2.3/antigravity-i18n.exe"><strong>下载 Windows x64 版</strong></a>
    ·
    <a href="https://github.com/wadewu-ml/regravity/releases">发布记录</a>
    ·
    已验证适配 <strong>Antigravity App 2.35.1</strong>
  </p>
  <p>
    <strong>简体中文</strong> · <a href="./README.md">English</a> · <a href="./README.zh-Hant.md">繁體中文</a> · <a href="./README.ja.md">日本語</a> · <a href="./README.ko.md">한국어</a> · <a href="./README.de.md">Deutsch</a> · <a href="./README.es.md">Español</a> · <a href="./README.fr.md">Français</a> · <a href="./README.pt-BR.md">Português (Brasil)</a> · <a href="./README.ru.md">Русский</a>
  </p>
  <p>antigravity-i18n 已更名为 Regravity。上方下载的是 v4.2.3。首个 Regravity 版本发布后，请手动下载 <code>regravity.exe</code>。新版会沿用现有补丁和自动更新设置。仓库更名后，v4.2.3 的内置检查可能提示无法比较版本号，可从上方「发布记录」查看更新。</p>
</div>

<p align="center">
  <a href="./docs/images/language-settings.webp">
    <img src="./docs/images/language-settings.webp" width="560" alt="Regravity 设置窗口，显示 Antigravity v2.35.2、已选简体中文，并同时开启会话指标与思考强度滑块">
  </a>
</p>

## 功能一览

- **单文件内置 9 种界面语言** — 翻译设置、对话区控件、模型菜单、原生菜单与退出提示，支持简体中文、繁体中文、日语、韩语、德语、西班牙语、法语、巴西葡萄牙语和俄语，不改动用户输入与对话内容。
- **可选会话指标与模型额度** — 在模型选择器右侧增加圆环，悬停即可查看模型剩余额度、本地重置时间、上下文已用、平均输出速度、缓存命中率与首 token 延迟。
- **可选思考强度滑块** — 在模型选择菜单中为支持调节的模型增加胶囊滑块，拖动即可切换思考强度。
- **独立组合，逐字节还原** — 两项开关可分别开启，也支持搭配官方英文界面；应用后即可关闭本工具，保留与当前 Antigravity 版本匹配的干净备份即可逐字节还原。

## 快速上手

### 应用设置

1. 双击运行 `antigravity-i18n.exe`。若未自动识别到正确目录，点 **更改位置** 选择包含 `Antigravity.exe` 的文件夹。
2. 在 **界面语言** 中选择目标语言（若只想保留英文界面并开启可选功能，选 **English** 即可）。
3. 按需开启语言列表下方的 **会话指标** 与 **思考强度滑块**。
4. 保存工作并退出 Antigravity，点击右下角 **应用更改**。Antigravity 会按新设置重新打开（也可点击 **打开 Antigravity** 启动），之后即可关闭本工具。

> [!WARNING]
> 若提示 Antigravity 仍在运行，请先保存工作后正常退出。使用 **强制退出并重试** 会中断进行中的会话，可能导致未保存的编辑丢失。

### 重新应用与还原

- **重新应用当前设置** — 点击右下角 **重新应用**。
- **还原官方版本** — 先退出 Antigravity，打开右上角菜单（`···`），点击 **还原官方版本** 并确认。界面将恢复官方英文，同时移除指标圆环和思考强度滑块。

> [!IMPORTANT]
> 首次应用时，工具会在 Antigravity 安装目录自动备份原文件至 `resources/app.asar.clean-backup`。如需日后还原，请保留该文件。

## 可选功能详解

### 会话指标与模型额度

- **发送首条消息前即可查额度** — 将鼠标移到模型选择器右侧的圆环上，即可查看各模型组在不同周期（如每周、5 小时）的剩余额度；停在百分比上可查看本地重置时间。
- **产生模型调用后显示实时指标** — 对话记录到模型调用后，圆环会反映上下文估算占比，面板上方同步显示 **上下文已用**（已用 token / 总容量与占比）、**平均输出速度**（含模型记录的思考 token）、**缓存命中率** 与 **首 token 延迟**。单击圆环可固定或收起面板。
- **多面板独立统计** — 分屏、侧问和子 Agent 面板各自读取本机对应会话最近一次有数据的模型调用。缺失项显示为 `—`。

<p align="center">
  <a href="./docs/images/session-quota-before-first-call.webp">
    <img src="./docs/images/session-quota-before-first-call.webp" width="440" alt="模型选择器右侧圆环上方展开的会话指标与模型剩余额度面板">
  </a>
</p>

### 思考强度滑块

在工具中开启 **思考强度滑块** 并应用更改。打开 Antigravity 模型选择菜单，支持调节思考强度的模型会在 **查看用量** 上方显示滑块——拖动档位预览，松开即生效。

<p align="center">
  <a href="./docs/images/thinking-effort-high.webp">
    <img src="./docs/images/thinking-effort-high.webp" height="340" alt="模型菜单中的思考强度滑块，当前为高">
  </a>
  &nbsp;&nbsp;
  <a href="./docs/images/thinking-effort-max.webp">
    <img src="./docs/images/thinking-effort-max.webp" height="340" alt="模型菜单中的思考强度滑块，当前为极高">
  </a>
</p>

## 更新说明

- **Antigravity 更新后** — 若翻译或增强组件失效，请先核对本页顶部的适配版本，再重新打开工具应用设置即可。
- **工具自动检查更新** — 启动时根据上次检查结果决定是否查询 GitHub。成功后至少间隔 24 小时再检查。右上角菜单（`···`）出现蓝点表示有新版本，点开可查看更新说明并就地下载重启。可在 **关于** 中关闭自动检查，或随时点击 **检查更新** 手动检查。工具更新完成后，点击 **重新应用** 即可将新译文写入 Antigravity。

## 命令行用法

在 `antigravity-i18n.exe` 所在目录打开 PowerShell：

| 操作 | 命令 |
| --- | --- |
| 列出语言代码 | `.\antigravity-i18n.exe locales` |
| 应用简体中文 | `.\antigravity-i18n.exe apply --locale zh-CN` |
| 官方英文 + 会话指标 | `.\antigravity-i18n.exe apply --locale en --session-metrics` |
| 官方英文 + 思考强度滑块 | `.\antigravity-i18n.exe apply --locale en --effort-slider` |
| 查看当前设置 | `.\antigravity-i18n.exe status` |
| 还原官方版本 | `.\antigravity-i18n.exe restore` |
| 查看全部参数 | `.\antigravity-i18n.exe --help` |

## 从源码构建

需要 Windows、PowerShell 7、Go 1.26.3+ 及 Wails v2 CLI：

```powershell
.\build.ps1 -Target gui
```

产物位于 `build/bin/regravity.exe`。

## 常见问题

- **启动时提示缺少 WebView2** — 请安装 [Microsoft Edge WebView2 运行时](https://developer.microsoft.com/en-us/microsoft-edge/webview2/#download) 后重新运行 `antigravity-i18n.exe`。

## 免责声明

1. 本项目为开源社区维护的非官方本地化工具，与 Google 或 Antigravity 官方无关。
2. Antigravity 及相关商标版权归原权利人所有。
3. 本工具仅供学习交流及个人使用，请在遵循相关软件使用协议的前提下合理使用。

[MIT 许可证](./LICENSE)
