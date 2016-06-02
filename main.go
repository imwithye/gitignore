package main

import (
	"fmt"
	"github.com/imwithye/gitignore/template"
	_ "github.com/imwithye/gitignore/template/all"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		ignore, _ := template.GetTemplate(arg)
		fmt.Println(ignore)
		fmt.Println()
	}
}
