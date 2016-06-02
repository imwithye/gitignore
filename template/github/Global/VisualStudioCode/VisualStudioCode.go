package VisualStudioCode

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		".vscode",
		"",
	}
	template.SetTemplate("GitHub/Global/VisualStudioCode", strings.Join(ignore, "\n"))
}
