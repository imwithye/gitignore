package fortran

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"C++.gitignore",
	}
	template.Add("github/fortran", strings.Join(ignore, "\n"))
}
