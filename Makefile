.phony: generate fmt build clean

build:
	go build -o git-ignore

generate:
	bash generate.sh
	go fmt ./...

fmt:
	go fmt ./...

clean:
	git clean -xdf
	rm git-ignore
