.DEFAULT_GOAL := build

TAG:=$(shell git rev-parse --short HEAD --dirty)

SRCS=$(wildcard *.go */*/*.go */*.go)

COMMANDS=$(wildcard cmd/*)
BIN_TARGETS=$(addprefix bin/, $(COMMANDS:cmd/%=%))

build: $(BIN_TARGETS)

$(BIN_TARGETS): $(SRCS)
	mkdir -p bin/
	go build -ldflags "-X main.buildref=$(TAG)" -o bin/$(@:bin/%=%) cmd/$(@:bin/%=%)/main.go

.PHONY: clean
clean:
	-rm $(BIN_TARGETS)

.PHONY: check
check:
	go test -race -cover -v ./...

.PHONY: lint
lint:
	which gometalinter > /dev/null || go get github.com/alecthomas/gometalinter
	gometalinter --install
	gometalinter --disable=gotype ./... --fast --deadline=10s
