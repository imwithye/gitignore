package commonlisp

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"*.FASL",
		"*.fasl",
		"*.lisp-temp",
		"",
	}
	template.Add("github/commonlisp", strings.Join(ignore, "\n"))
}
