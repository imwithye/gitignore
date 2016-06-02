.phony: generate fmt build clean

build:
	env GOOS=darwin GOARCH=amd64 go build -o git-ignore_darwin_amd64
	env GOOS=windows GOARCH=amd64 go build -o git-ignore_windows_amd64
	env GOOS=windows GOARCH=386 go build -o git-ignore_windows_386
	env GOOS=linux GOARCH=amd64 go build -o git-ignore_linux_amd64
	env GOOS=linux GOARCH=386 go build -o git-ignore_linux_386
	env GOOS=linux GOARCH=arm go build -o git-ignore_linux_arm

generate:
	bash generate.sh
	go fmt ./...

fmt:
	go fmt ./...

clean:
	rm *.zip
	rm -rf gitignore-master
	rm git-ignore_*
