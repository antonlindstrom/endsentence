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
	golangci-lint --version &> /dev/null || GO111MODULE=on go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.17.1
	golangci-lint run --enable-all -D gochecknoglobals -D typecheck
