package Coq

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"*.vo",
		"*.glob",
		"*.v.d",
	}
	template.SetTemplate("GitHub/Coq", strings.Join(ignore, "\n"))
}
