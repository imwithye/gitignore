package main

import (
	"embed"
	"fmt"
	"os"
	"regexp"
	"strings"
)

//go:embed templates
var templates embed.FS

func gitignore(args []string) string {
	var ret []string
	for _, arg := range args {
		ret = append(ret, fmt.Sprintf("# Gitignore %s", arg))
		data, err := templates.ReadFile(fmt.Sprintf("templates/%v.gitignore", arg))
		if err != nil {
			continue
		}
		ret = append(ret, string(data))
		ret = append(ret, "\n")
	}
	str := strings.Join(ret, "\n")
	str = strings.ReplaceAll(str, "\r", "\n")
	reg := regexp.MustCompile(`([\n]{3})+`)
	str = reg.ReplaceAllString(str, "\n\n")
	str = strings.ReplaceAll(str, "\n\n", "\n")
	reg = regexp.MustCompile(`\n*$`)
	str = reg.ReplaceAllString(str, "\n")
	return str
}

func main() {
	args := os.Args[1:]
	fmt.Println(gitignore(args))
}
