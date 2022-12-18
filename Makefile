default: build

build:
	go build ./cmd/xoq

docker:
	docker build -t xoq-backend -f build/package/Dockerfile .
	docker build -t xoq-frontend web

run:
	go run ./cmd/xoq

hooks:
	cp githooks/* .git/hooks/

.PHONY: build docker run hooks
