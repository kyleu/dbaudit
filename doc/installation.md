<!--- Content managed by Project Forge, see [projectforge.md] for details. -->
# Installation

## Pre-built binaries
Download any package from the [release page](https://github.com/kyleu/dbaudit/releases).

## Running with Docker
```shell
docker run -p 55500:55500 ghcr.io/kyleu/dbaudit:latest
docker run -p 55500:55500 ghcr.io/kyleu/dbaudit:latest-debug
```

## Built from source

### go install
```shell
go install github.com/kyleu/dbaudit@latest
```

### Source code

If you want to contribute to the project, please follow the steps on our [contributing guide](contributing).

If you just want to build from source for whatever reason, follow these steps:

```shell
git clone https://github.com/kyleu/dbaudit
cd dbaudit
make build
./build/debug/dbaudit --help
```

A script has been provided at `./bin/dev.sh` that will auto-reload when the project changes.

Note that you may need to run `./bin/bootstrap.sh` before building the project for the first time.
