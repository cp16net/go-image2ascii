default: all
all: generate build run

generate:
	govendor generate

build:
	govendor build

run:
	./go-image2ascii http

clean:
	rm -f go-image2ascii
