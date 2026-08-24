# Electerm 同步服务器 Go 版本

[![English](https://img.shields.io/badge/Language-English-blue.svg)](README.md)
[![中文](https://img.shields.io/badge/语言-中文-red.svg)](README_CN.md)

开源终端/ssh/telnet/serialport/RDP/VNC/Spice/sftp/ftp客户端(Linux, Mac, Windows, Android, HarmonyOS)。

除了主流的 Windows/macOS/Linux/Android，electerm 还支持鸿蒙(HarmonyOS)，以及老旧系统——如 Ubuntu 18、Windows 7、macOS 10+，以及国产特殊 Linux 发行版如 UOS、麒麟(Kylin)、龙芯(LoongArch，含旧世界与新世界)。

<p>
  <a href="https://electerm.org">主页 / 下载</a> ·
  <a href="https://theme.electerm.org">主题</a> ·
  <a href="https://github.com/electerm/electerm-web-docker">Docker</a> ·
  <a href="https://demo.electerm.org">在线演示</a> ·
  <a href="https://github.com/electerm/electerm-android">Android</a> ·
  <a href="https://github.com/electerm/electerm-harmony">鸿蒙</a> ·
  <a href="https://appgallery.huawei.com/app/detail?id=org.electerm.electerm">华为应用市场</a> ·
  <a href="https://www.microsoft.com/store/apps/9NCN7272GTFF">微软商店</a> ·
  <a href="https://snapcraft.io/electerm">Snap 商店</a> ·
  <a href="https://repos.electerm.org/deb">deb 仓库</a> ·
  <a href="https://repos.electerm.org/rpm">rpm 仓库</a>
</p>

<div>🌐 <strong><a href="https://cloud.electerm.org">electerm 在线版</a></strong> — 公共免费在线 electerm 应用</div>
<div>🤖 <strong><a href="https://ai.electerm.org">electerm AI</a></strong> — 免费为 electerm 用户提供 AI</div>
<div>💻 <strong><a href="https://github.com/electerm/electerm-web">electerm-web</a></strong> — 运行于浏览器(支持移动设备)的 web app 版本</div>

一个简单的 Electerm 数据同步服务器，使用 Go 语言编写。

## 使用方法

需要 Go 1.16+

```bash
git clone git@github.com:electerm/electerm-sync-server-go.git
cd electerm-sync-server-go

# 安装依赖
go mod download

# 创建环境配置文件，然后编辑 .env
cp sample.env .env

# 开发模式运行
go run src/main.go

# 会显示类似信息：
# server running at http://127.0.0.1:7837

# 在 Electerm 同步设置中，设置自定义同步服务器：
# 服务器 URL: http://127.0.0.1:7837
# 然后可以在 Electerm 自定义同步中使用 http://127.0.0.1:7837/api/sync 作为 API URL

# JWT_SECRET: .env 文件中的 JWT_SECRET
# JWT_USER_NAME: .env 文件中的一个 JWT_USER
```

## 生产环境构建和运行

对于类 Unix 系统（Linux/macOS）：

```bash
# 运行构建脚本
./bin/build.sh

# 配置 .env 后运行服务器
GIN_MODE=release ./output/electerm-sync-server-go
```

## 测试

```bash
bin/test.sh
```

## 编写自己的数据存储

以 [src/store/sql.go](src/store/sql.go) 为例，编写自己的读写方法。默认存储现在使用 SQLite 以获得更好的性能和可靠性。

## 其他语言的同步服务器

[https://github.com/electerm/electerm/wiki/Custom-sync-server](https://github.com/electerm/electerm/wiki/Custom-sync-server)

## 许可证

MIT
