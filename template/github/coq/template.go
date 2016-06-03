package coq

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"*.vo",
		"*.glob",
		"*.v.d",
		"",
	}
	template.Add("github/coq", strings.Join(ignore, "\n"))
}
