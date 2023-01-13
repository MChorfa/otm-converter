---
# POC - DO NOT USE
---

# otm-converter

Attempt to build a converter to generate an open threat model based on the the work from [IriusRisk](https://github.com/iriusrisk)

- Specification: https://github.com/iriusrisk/OpenThreatModel
- Inspiration: https://github.com/iriusrisk/startleft

---
# POC - DO NOT USE
---


## Init Project

```shell
# initialize the project
go mod init github.com/MChorfa/otm-converter
go mod tidy
# open project in VSCode
code .
```



## Build

```shell
# print the version of conveyor
go run main.go --version
conveyor version v0.0.1-alpha

## 
# build the converter CLI in version v0.0.1-alpha
go build -o ./dist/converter -ldflags="-X 'github.com/MChorfa/otm-converter/cmd/converter.version=v0.0.1-alpha'" main.go

# verify version is being set correctly
./dist/converter --version
> converter version v0.0.1-alpha
```


## Run

```shell
# converter
go run main.go convert \
--input-path "./test/integration/data/design.json" \
--output-path "./test/integration/data"
```

# Release

```sh
# https://goreleaser.com/quick-start/

brew install goreleaser/tap/goreleaser
goreleaser init
goreleaser build --single-target --snapshot --rm-dist
goreleaser release --snapshot --rm-dist

# The minimum permissions the GITHUB_TOKEN should have to run this are write:packages
export GITHUB_TOKEN="YOUR_GH_TOKEN"
git tag -a v0.0.1 -m "Pre release"
git push origin v0.1.0
goreleaser release
```