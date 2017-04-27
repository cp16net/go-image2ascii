default: all
all: generate build run

generate:
	govendor generate

build:
	govendor build

run:
	PORT=8888 ./go-image2ascii

clean:
	rm -f go-image2ascii
