default: build

build:
	go build ./cmd/xoq

run:
	go run ./cmd/xoq

hooks:
	cp githooks/* .git/hooks/

.PHONY: build run hooks
