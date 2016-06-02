.phony: generate fmt build clean

build:
	go build -o git-ignore

generate:
	bash generate.sh
	go fmt ./...

fmt:
	go fmt ./...

clean:
	rm *.zip
	rm -rf gitignore-master
	rm git-ignore
