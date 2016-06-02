package main

import (
	"fmt"
	"github.com/imwithye/git_ignore/template"
	_ "github.com/imwithye/git_ignore/template/all"
)

func main() {
	fmt.Println(template.GetTemplate("GitHub/C"))
}
