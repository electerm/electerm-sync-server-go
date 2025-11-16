# Electerm 同步服务器 Go 版本

[![English](https://img.shields.io/badge/Language-English-blue.svg)](README.md)
[![中文](https://img.shields.io/badge/语言-中文-red.svg)](README_CN.md)

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
