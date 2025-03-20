GO_BIN ?= go
OUT_BIN = main

export PATH := $(PATH):/usr/local/go/bin

build:
	$(GO_BIN) mod tidy
	$(GO_BIN) build -o $(OUT_BIN) -v

download:
	$(GO_BIN) get
	$(GO_BIN) mod tidy

update:
	$(GO_BIN) get -u
	$(GO_BIN) mod tidy
