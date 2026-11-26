<div align="center">
  <img src="./frontend/dist/appicon.png" width="68" height="68" alt="Regravity 图标">
  <h1>Regravity</h1>
  <p><em>反重力起飞，Regravity 让它落地——回归你的语言与掌控。</em></p>
  <p>专为 Windows 版 <strong>Google Antigravity</strong> 打造的便携增强工具：15 种界面语言、会话指标与模型额度，以及思考强度滑块。</p>
  <p>
    <a href="https://github.com/wadewu-ml/regravity/releases/latest/download/regravity.exe"><img src="https://img.shields.io/badge/Download-Latest_(Windows_x64)-0F9D58?style=flat&logo=windows&logoColor=white" alt="下载最新版 Windows x64 Regravity"></a>
    <a href="https://github.com/wadewu-ml/regravity/releases"><img src="https://img.shields.io/badge/Releases-GitHub-18181B?style=flat&logo=github" alt="GitHub 发布记录"></a>
    <img src="https://img.shields.io/badge/Antigravity-2.35.2_tested-18181B?style=flat" alt="已验证适配 Antigravity App 2.35.2">
    <a href="./LICENSE"><img src="https://img.shields.io/badge/License-MIT-18181B?style=flat" alt="MIT 许可证"></a>
  </p>
  <p>
    <sub>
      <strong>简体中文</strong> · <a href="./README.md">English</a> · <a href="./README.zh-Hant.md">繁體中文</a> · <a href="./README.ja.md">日本語</a> · <a href="./README.ko.md">한국어</a> · <a href="./README.de.md">Deutsch</a> · <a href="./README.es.md">Español</a> · <a href="./README.fr.md">Français</a> · <a href="./README.pt-BR.md">Português (Brasil)</a> · <a href="./README.ru.md">Русский</a> · <a href="./README.tr.md">Türkçe</a> · <a href="./README.vi.md">Tiếng Việt</a> · <a href="./README.it.md">Italiano</a> · <a href="./README.ar.md">العربية</a> · <a href="./README.id.md">Bahasa Indonesia</a> · <a href="./README.pl.md">Polski</a>
    </sub>
  </p>
</div>

<br>

<p align="center">
  <a href="./docs/images/language-settings.webp">
    <img src="./docs/images/language-settings.webp" width="580" alt="Regravity 设置窗口，显示 Antigravity v2.35.2、已选简体中文，并同时开启会话指标与思考强度滑块">
  </a>
  <br>
  <sub>Regravity 设置窗口，图中客户端为 Antigravity 2.35.2。</sub>
</p>

## 核心能力

- **15 种界面语言，不碰对话内容** — 翻译设置、对话区控件、模型菜单、原生菜单与退出提示，完全不改动用户输入与聊天记录。
- **会话指标与模型额度** *（可选）* — 在模型选择器右侧增加圆环，悬停即看各模型组剩余额度、本地重置时间、上下文已用、平均输出速度、缓存命中率与首 token 延迟。
- **思考强度滑块** *（可选）* — 在模型选择菜单中为支持调节的模型嵌入胶囊滑块，拖动即切档位。
- **独立组合，逐字节还原** — 两项增强开关可单独启用，也兼容官方英文界面；应用后无需常驻后台，保留与当前 Antigravity 版本匹配的干净备份即可逐字节还原。

## 快速上手

在右上角可以选择设置窗口的语言。

### 应用设置

1. 双击运行 `regravity.exe`。若未自动识别正确目录，点 **更改位置** 选择包含 `Antigravity.exe` 的文件夹。
2. 在 **界面语言** 中选择目标语言（只想保留英文界面并开启可选功能时，选 **English** 即可）。
3. 按需开启语言列表下方的 **会话指标** 与 **思考强度滑块**。
4. 保存工作并退出 Antigravity，点击右下角 **应用更改**。Antigravity 会按新设置重新打开（也可点 **打开 Antigravity** 启动），之后即可关闭本工具。

> [!WARNING]
> 若提示 Antigravity 仍在运行，请先保存工作并正常退出。使用 **强制退出并重试** 会中断进行中的会话，可能导致未保存的编辑丢失。

### 重新应用与还原

- **重新应用** — 更新工具后，点一下 **重新应用** 就能用上新译文。
- **还原官方版本** — 先退出 Antigravity，打开右上角菜单（`···`），点击 **还原官方版本** 并确认，界面将恢复官方英文并移除指标圆环与滑块。

> [!IMPORTANT]
> 首次应用时，工具会在 Antigravity 安装目录自动备份原文件至 `resources/app.asar.clean-backup`。如需日后还原，请保留该文件。

## 可选功能

### 会话指标与模型额度

- **发送首条消息前即可查额度** — 悬停在模型选择器右侧圆环上，即可查看各模型组在不同周期（如每周、5 小时）的剩余额度；停在百分比上可查看本地重置时间。
- **产生模型调用后显示实时指标** — 对话记录到模型调用后，圆环反映估算上下文占比，面板上方自动填充 **上下文已用**（已用 token / 总容量与占比）、**平均输出速度**（含模型记录的思考 token）、**缓存命中率** 与 **首 token 延迟**。单击圆环可固定或收起。
- **多面板独立统计** — 分屏、侧问和子 Agent 面板各自读取对应会话最近一次有数据的本地记录，缺失项显示为 `—`。

<p align="center">
  <a href="./docs/images/session-quota-before-first-call.webp">
    <img src="./docs/images/session-quota-before-first-call.webp" width="420" alt="模型选择器右侧圆环上方展开的会话指标与模型剩余额度面板">
  </a>
  <br>
  <sub>新对话即可直接查看各模型组剩余额度；产生首次调用后上方指标自动填充。</sub>
</p>

### 思考强度滑块

开启 **思考强度滑块** 并应用更改后，打开 Antigravity 模型选择菜单，支持直接调节思考强度的模型会在 **查看用量** 上方显示滑块——拖动档位预览，松手即生效。

<p align="center">
  <a href="./docs/images/thinking-effort-high.webp">
    <img src="./docs/images/thinking-effort-high.webp" height="340" alt="模型菜单中的思考强度滑块，当前为高">
  </a>
  &nbsp;&nbsp;
  <a href="./docs/images/thinking-effort-max.webp">
    <img src="./docs/images/thinking-effort-max.webp" height="340" alt="模型菜单中的思考强度滑块，当前为极高">
  </a>
  <br>
  <sub>常规档位（高）与高亮最高档位（极高），直接在模型菜单中拖动切换。</sub>
</p>

## 更新与命令行

### 更新说明

- **Antigravity 更新后** — 若翻译或增强组件消失，请先核对顶部已验证的客户端版本，再重新打开工具应用设置即可。
- **工具自动检查更新** — 启动时根据上次检查结果决定是否查询 GitHub，成功后至少间隔 24 小时再查。右上角菜单（`···`）出现蓝点表示有新版本，点开可查看更新说明并就地更新重启。可在 **关于** 中关闭自动检查，或随时点击 **检查更新** 手动检查。更新完成后点击 **重新应用** 即可写入新译文。

如果你用的是 v4.2.3 及更早版本，请手动下载并运行 [`regravity.exe`](https://github.com/wadewu-ml/regravity/releases/latest/download/regravity.exe)，原有设置会保留。

### 命令行用法

在 `regravity.exe` 所在目录打开 PowerShell：

| 操作 | 命令 |
| --- | --- |
| 列出语言代码 | `.\regravity.exe locales` |
| 应用简体中文 | `.\regravity.exe apply --locale zh-CN` |
| 官方英文 + 会话指标 | `.\regravity.exe apply --locale en --session-metrics` |
| 官方英文 + 思考强度滑块 | `.\regravity.exe apply --locale en --effort-slider` |
| 查看当前设置 | `.\regravity.exe status` |
| 还原官方版本 | `.\regravity.exe restore` |
| 查看全部参数 | `.\regravity.exe --help` |

### 从源码构建

需要 Windows、PowerShell 7、Go 1.26.3+ 及 Wails v2 CLI：

```powershell
.\build.ps1 -Target gui
```

产物输出至 `build/bin/regravity.exe`。

## 常见问题

- **启动时提示缺少 WebView2** — 请安装 [Microsoft Edge WebView2 运行时](https://developer.microsoft.com/en-us/microsoft-edge/webview2/#download) 后重新运行 `regravity.exe`。

---

## 免责声明

1. 本项目为开源社区维护的非官方本地化工具，与 Google 或 Antigravity 官方无关。
2. Antigravity 及相关商标版权归原权利人所有。
3. 本工具仅供学习交流及个人使用，请在遵循相关软件使用协议的前提下合理使用。

[MIT 许可证](./LICENSE)
