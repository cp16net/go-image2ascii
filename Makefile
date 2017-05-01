FILES=$(shell go list ./... | grep -v /vendor/ | grep -v /gen/)

default: help

help:
	@echo "These 'make' targets are available."
	@echo
	@echo "  all                Clean, generate, test, build, vet, install"
	@echo "  clean              Removes all build output"
	@echo "  generate           Generates all code needed"
	@echo "  test               Run the unit tests"
	@echo "  build              Builds the binary"
	@echo "  vet                Runs govendor vet against proejct"
	@echo "  install            Installs the binary to ${GOBIN}"
	@echo

all: clean generate test build vet install

clean:
	rm -f go-image2ascii
	rm -f ${GOBIN}/go-image2ascii
	rm -rf gen

generate:
	govendor generate
	swagger generate server -t gen -f ./swagger/swagger.yml --exclude-main -A converter

test:
	govendor test -cover $(FILES)

build:
	govendor build

vet:
	govendor vet $(FILES)

install:
	govendor install
