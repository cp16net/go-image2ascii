default: all
all: generate build vet run

generate:
	govendor generate

build:
	govendor build

vet:
	govendor vet

test:
	go test ./... | grep -v /vendor/

run:
	./go-image2ascii iris

install:
	govendor install

clean:
	rm -f go-image2ascii
	rm -f ${GOBIN}/go-image2ascii
