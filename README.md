<h1 align="center" style="padding-top: 60px;padding-bottom: 40px;">
    <a href="https://electerm.org">
        <img src="https://github.com/electerm/electerm-resource/raw/master/static/images/electerm.png", alt="" />
    </a>
</h1>

[![English](https://img.shields.io/badge/Language-English-blue.svg)](README.md)
[![中文](https://img.shields.io/badge/语言-中文-red.svg)](README_CN.md)

# Electerm sync server go


A simple electerm data sync server with go.

## Use

Requires go 1.16+

```bash
git clone git@github.com:electerm/electerm-sync-server-go.git
cd electerm-sync-server-go

# Install dependencies
go mod download

# create env file, then edit .env
cp sample.env .env

# Run in development mode
go run src/main.go

# would show something like
# server running at http://127.0.0.1:7837

# in electerm sync settings, set custom sync server with:
# server url: http://127.0.0.1:7837
# Then you can use http://127.0.0.1:7837/api/sync as API Url in electerm custom sync

# JWT_SECRET: your JWT_SECRET in .env
# JWT_USER_NAME: one JWT_USER in .env
```

## Build and Run in production

For Unix-like systems (Linux/macOS):

```bash
# Run the build script
./bin/build.sh

# Run the server (after configuring .env)
GIN_MODE=release ./output/electerm-sync-server-go
```

## Test

```bash
bin/test.sh
```

## Write your own data store

Just take [src/store/sql.go](src/store/sql.go) as an example, write your own read/write method. The default storage is now SQLite for better performance and reliability.

## Sync server in other languages

[https://github.com/electerm/electerm/wiki/Custom-sync-server](https://github.com/electerm/electerm/wiki/Custom-sync-server)

## About electerm

Open-sourced terminal/ssh/sftp/telnet/serialport/RDP/VNC/Spice/ftp client(Linux, Mac, Windows, Android, HarmonyOS).

Besides mainstream Windows/macOS/Linux/Android, electerm also supports HarmonyOS, and older systems — Ubuntu 18, Windows 7, macOS 10+, and special Chinese Linux distributions such as UOS, Kylin, and LoongArch (both old-world and new-world).

<p>
  <a href="https://electerm.org">Homepage / Downloads</a> ·
  <a href="https://theme.electerm.org">Theme</a> ·
  <a href="https://github.com/electerm/electerm-web-docker">Docker</a> ·
  <a href="https://demo.electerm.org">Online demo</a> ·
  <a href="https://github.com/electerm/electerm-android">Android</a> ·
  <a href="https://github.com/electerm/electerm-harmony">HarmonyOS</a> ·
  <a href="https://appgallery.huawei.com/app/detail?id=org.electerm.electerm">Huawei AppGallery</a> ·
  <a href="https://www.microsoft.com/store/apps/9NCN7272GTFF">Microsoft Store</a> ·
  <a href="https://snapcraft.io/electerm">Snap Store</a> ·
  <a href="https://repos.electerm.org/deb">deb repo</a> ·
  <a href="https://repos.electerm.org/rpm">rpm repo</a>
</p>

<div>🌐 <strong><a href="https://cloud.electerm.org">electerm online</a></strong> — Public free online electerm app</div>
<div>🤖 <strong><a href="https://ai.electerm.org">electerm AI</a></strong> — Free AI for electerm users</div>
<div>💻 <strong><a href="https://github.com/electerm/electerm-web">electerm-web</a></strong> — Web app version running in browser (including mobile device)</div>

## License

MIT
