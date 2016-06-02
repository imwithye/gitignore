package main

import (
	"fmt"
	"github.com/imwithye/gitignore/template"
	_ "github.com/imwithye/gitignore/template/all"
	"os"
)

const help = `usage: git ignore <ignores>
	-h	--help			print this help message
	-v	--version		print version information
`
const version = "v0.1.0"

func main() {
	if len(os.Args[1:]) == 0 || os.Args[1:][0] == "-h" || os.Args[1:][0] == "--help" {
		fmt.Println(help)
		return
	}
	if os.Args[1:][0] == "-v" || os.Args[1:][0] == "--version" {
		fmt.Println(version)
		return
	}
	for _, arg := range os.Args[1:] {
		ignore, _ := template.GetTemplate(arg)
		fmt.Println(ignore)
		fmt.Println()
	}
}
