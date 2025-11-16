# Electerm sync server go

[![English](https://img.shields.io/badge/Language-English-blue.svg)](README.md)
[![中文](https://img.shields.io/badge/语言-中文-red.svg)](README_CN.md)

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

## License

MIT
