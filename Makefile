export CGO_ENABLED=0
export GOMOD111MODULE=on

dist: build

build:
	go build -v -o bin/es-cleaner cmd/es-cleaner/main.go
