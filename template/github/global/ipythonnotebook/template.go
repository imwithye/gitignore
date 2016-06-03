package ipythonnotebook

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# Temporary data",
		".ipynb_checkpoints/",
		"",
	}
	template.Add("github/global/ipythonnotebook", strings.Join(ignore, "\n"))
}
