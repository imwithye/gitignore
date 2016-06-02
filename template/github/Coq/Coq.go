package Coq

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"*.vo",
		"*.glob",
		"*.v.d",
	}
	template.SetTemplate("GitHub/Coq", strings.Join(ignore, "\n"))
}
