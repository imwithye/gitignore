.phony: generate fmt build

build:
	go build -o git-ignore

generate:
	bash generate.sh
	go fmt ./...

fmt:
	go fmt ./...

clean:
	rm git-ignore
