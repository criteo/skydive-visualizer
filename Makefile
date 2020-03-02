GOOS ?= linux
GOARCH ?= amd64
OUT ?= skydive-visualizer

.PHONY: static deps front-deps front

deps:
	go get github.com/rakyll/statik

static: deps front
	statik -f -src=./frontend/build -dest=server/ -p public

release: static
	env GOOS=$(GOOS) GOARCH=$(GOARCH) go build -tags release -o $(OUT)

front-deps:
	cd frontend && yarn install

front: front-deps
	cd frontend && yarn build

