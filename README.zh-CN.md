# <img src="./frontend/dist/appicon.png" width="32" height="32" alt=""> antigravity-i18n

为 Windows 版 Google Antigravity 提供九种界面语言包，翻译设置、对话区按钮、菜单和退出提示。还可以在输入框旁查看上下文余量和模型的输出速度。设置生效后即可关闭本工具。

[下载 Windows x64 版](https://github.com/wadewu-ml/antigravity-i18n/releases/latest/download/antigravity-i18n.exe) · [发布记录](https://github.com/wadewu-ml/antigravity-i18n/releases)

适配 **Antigravity App 2.35.1**。

**简体中文** · [English](./README.md) · [繁體中文](./README.zh-Hant.md) · [日本語](./README.ja.md) · [한국어](./README.ko.md) · [Deutsch](./README.de.md) · [Español](./README.es.md) · [Français](./README.fr.md) · [Português (Brasil)](./README.pt-BR.md) · [Русский](./README.ru.md)

<a href="./docs/images/antigravity-zh-cn.png"><img src="./docs/images/translated-window.webp" width="680" alt="简体中文界面，顶部是中文菜单，输入框里可以看到模型选择器和指标圆环"></a>

## 切换界面语言

可选简体中文、繁体中文、日语、韩语、德语、西班牙语、法语、巴西葡萄牙语和俄语。

1. 双击 `antigravity-i18n.exe`。如果工具没找到正确的安装目录，点「更改位置」，选择包含 `Antigravity.exe` 的文件夹。
2. 在「界面语言」中选好语言，按需打开「会话指标」。想保留英文界面并查看指标，语言选 **English** 即可。
3. 保存工作并退出 Antigravity，再点「应用更改」。Antigravity 会按新设置重新打开。

退出提示中的「强制退出并重试」会中断进行中的会话，未保存的编辑可能丢失。

<a href="./docs/images/language-settings.webp"><img src="./docs/images/language-settings.webp" width="640" alt="工具中已选简体中文，并已打开「会话指标」开关"></a>


## 查看会话指标

对话有了模型调用记录后，输入框里的模型选择器右侧会出现一个圆环，表示上下文大约用了多少。将鼠标移到圆环上，可以查看剩余 token 数、输出速度、缓存命中率和首个 token 的等待时间。单击圆环可固定详情，再点一次收起。

<a href="./docs/images/session-metrics-zh-cn.png"><img src="./docs/images/session-metrics-detail.webp" width="560" alt="模型选择器右侧的圆环上方展开了会话指标，显示上下文用量、余量和输出速度"></a>

指标取自本机的会话记录。分屏、侧问和子 Agent 面板各自显示对应会话最近一次有数据的模型调用。输出量包含模型记录的思考 token。`—` 表示这一项暂时没有数据。

## 还原与更新

要恢复官方英文界面，先退出 Antigravity，在工具里点「还原官方版本」，再点「还原」。会话指标也会随之关闭。首次应用时工具会自动备份原文件，请保留安装目录下的 `resources/app.asar.clean-backup`。

如果 Antigravity 更新后翻译或指标消失，可先对照本页顶部的适配版本，再选择语言和指标并应用。要再次应用已经生效的设置，点「重新应用」。

要更新工具，点窗口顶部版本号旁的「检查更新」。工具会在这时查询 GitHub。点「更新并重启」后，新版会从原来的 EXE 路径打开。需要把新译文应用到 Antigravity 时，再点「重新应用」。

## 命令行用法

在 `antigravity-i18n.exe` 所在目录打开 PowerShell。

| 操作 | 命令 |
| --- | --- |
| 列出语言代码 | `.\antigravity-i18n.exe locales` |
| 应用简体中文 | `.\antigravity-i18n.exe apply --locale zh-CN` |
| 使用官方英文并开启指标 | `.\antigravity-i18n.exe apply --locale en --session-metrics` |
| 查看当前设置 | `.\antigravity-i18n.exe status` |
| 还原官方版本 | `.\antigravity-i18n.exe restore` |
| 查看全部参数 | `.\antigravity-i18n.exe --help` |

## 从源码构建

在 Windows 上安装 PowerShell 7、Go 1.26.3 或更新版本，以及 Wails v2 CLI。在仓库根目录运行 `.\build.ps1 -Target gui`，产物位于 `build/bin/antigravity-i18n.exe`。

## 启动问题

如果启动工具时提示缺少 WebView2，请安装 [Microsoft Edge WebView2 运行时](https://developer.microsoft.com/en-us/microsoft-edge/webview2/#download)。

## 免责声明

1. 本项目为开源社区维护的非官方本地化工具，与 Google 或 Antigravity 官方无关。
2. Antigravity 及相关商标版权归原权利人所有。
3. 本工具仅供学习交流及个人使用，请在遵循相关软件使用协议的前提下合理使用。

[MIT 许可证](./LICENSE)
