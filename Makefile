.PHONY: all build lint test benchmark gocyclo

PROJECT_NAME := __name_your_project__
VERSION := $(shell git describe --tags --always --dirty="-dev")

all: clean build

clean:
	rm -rf ${PROJECT_NAME} build/
	mkdir build
	touch build/.gitkeep

build:
	go build -trimpath -ldflags "-X main.version=${VERSION}" -v -o build/${PROJECT_NAME} main.go

lint:
	gofmt -d -s ./
	go vet ./...
	staticcheck ./...

test:
	go test ./...

benchmark:
	go test ./... -bench=. -benchmem