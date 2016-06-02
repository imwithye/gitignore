package VisualStudioCode

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		".vscode",
		"",
	}
	template.SetTemplate("GitHub/Global/VisualStudioCode", strings.Join(ignore, "\n"))
}
