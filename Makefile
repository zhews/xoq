default: build

build:
	go build ./cmd/xoq

docker:
	docker build -t xoq -f build/package/xoq/Dockerfile .

run:
	go run ./cmd/xoq

hooks:
	cp githooks/* .git/hooks/

.PHONY: build run hooks
