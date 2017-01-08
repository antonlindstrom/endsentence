.DEFAULT_GOAL := build

TAG:=$(shell git rev-parse --short HEAD --dirty)

build: endsentence

endsentence: main.go linter/linter.go
	go build -ldflags "-X main.buildref=$(TAG)"

.PHONY: clean
clean:
	-rm -r endsentence

.PHONY: check
check:
	go test -race -cover -v ./linter/...

.PHONY: lint
lint:
	which gometalinter > /dev/null || go get github.com/alecthomas/gometalinter
	gometalinter --install
	gometalinter --disable=gotype ./... --fast --deadline=10s
