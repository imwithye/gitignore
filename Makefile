.PHONY: build build_all generate fmt clean

build:
	mkdir -p build
	go build -o build/git-ignore

build_all:
	mkdir -p build
	env GOOS=darwin GOARCH=amd64 go build -o build/git-ignore_darwin_amd64
	env GOOS=windows GOARCH=amd64 go build -o build/git-ignore_windows_amd64.exe
	env GOOS=windows GOARCH=386 go build -o build/git-ignore_windows_386.exe
	env GOOS=linux GOARCH=amd64 go build -o build/git-ignore_linux_amd64
	env GOOS=linux GOARCH=386 go build -o build/git-ignore_linux_386
	env GOOS=linux GOARCH=arm go build -o build/git-ignore_linux_arm

generate:
	go run generator/generate.go
	make fmt

fmt:
	go fmt ./...

clean:
	rm -rf *.zip git-ignore
	rm -rf gitignore-master build
