# 打包与分发

本项目只分发源码归档和本地 `.tgz` 安装包，不向 npm 注册表发布。`package.json` 设置 `private: true`，阻止意外执行 `npm publish`；`npm pack` 仍可本地打包。

在项目根目录执行：

```powershell
npm ci
New-Item -ItemType Directory -Force dist
npm pack --pack-destination dist
```

`npm pack` 通过 `prepack` 自动运行测试。确认成功后，可分发 `dist/` 中的包，或将包作为 GitHub Release 附件。源码归档可使用 GitHub 的源码下载，或对已提交的版本执行 `git archive`。不要移动既有版本标签；正式新版本应使用新版本号。

解压 `.tgz` 后进入 `package` 目录，安装运行依赖，再运行：

```powershell
npm install --omit=dev
node bin/cli.js --version
node bin/cli.js apply --locale zh-CN
```

安装依赖可能联网；这与向 npm 上传发布项目不同。包内不包含 `node_modules`。已有 npm 版本不在此流程中删除或更新。
