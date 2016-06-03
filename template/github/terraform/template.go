package terraform

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# Compiled files",
		"*.tfstate",
		"*.tfstate.backup",
		"",
	}
	template.Add("github/terraform", strings.Join(ignore, "\n"))
}
