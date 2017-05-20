FILES=$(shell go list ./... | grep -v /vendor/ | grep -v /gen/)

default: help

help:
	@echo "These 'make' targets are available."
	@echo
	@echo "  all                Clean, generate, test, build, vet, install"
	@echo "  clean              Removes all build output"
	@echo "  generate           Generates all code needed"
	@echo "  test               Run the unit tests"
	@echo "  vet                Runs govendor vet against proejct"
	@echo "  build              Builds the binary for all GOOS=linux/darwin/windows"
	@echo "  install            Installs the binary to ${GOBIN}"
	@echo

all: clean generate test vet build install
	@echo "success - everything is awesome"

release: clean generate test vet build-all

clean:
	rm -f go-image2ascii
	rm -f ${GOBIN}/go-image2ascii
	rm -rf gen dist

generate:
	govendor generate
	swagger generate server -t gen -f ./swagger.yml --exclude-main -A converter

test:
	govendor test -cover $(FILES)

build:
	mkdir -p dist
	govendor build -o dist/go-image2ascii .

build-all: build-linux build-darwin build-windows

build-linux:
	mkdir -p dist
	GOOS=linux govendor build -o dist/go-image2ascii-linux .

build-darwin:
	mkdir -p dist
	GOOS=darwin govendor build -o dist/go-image2ascii-darwin .

build-windows:
	mkdir -p dist
	GOOS=windows govendor build -o dist/go-image2ascii-win.exe .

vet:
	govendor vet $(FILES)

install:
	govendor install

http:
	go-image2ascii http
