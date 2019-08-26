export CGO_ENABLED=0
export GOMOD111MODULE=on

dist: build

build:
	go build -v -o bin/es-cleaner cmd/es-cleaner/main.go
	go build -v -o bin/test cmd/test/main.go

test:
	go test -v ./...

-include local.mk
