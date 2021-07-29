package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"
)

//go:embed templates
var templates embed.FS

func main() {
	args := os.Args[1:]
	var ret []string
	for _, arg := range args {
		data, err := templates.ReadFile(fmt.Sprintf("templates/%v.gitignore", arg))
		if err != nil {
			continue
		}
		ret = append(ret, string(data))
	}
	ioutil.WriteFile(".gitignore", []byte(strings.Join(ret, "\n")), fs.ModeAppend)
}
