package Xojo

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"# Xojo (formerly REALbasic and Real Studio)",
		"",
		"Builds*",
		"*.debug",
		"*.debug.app",
		"Debug*.exe",
		"Debug*/Debug*.exe",
		"Debug*/Debug*\\ Libs",
		"*.rbuistate",
		"*.xojo_uistate",
		"*.obsolete",
	}
	template.SetTemplate("GitHub/Xojo", strings.Join(ignore, "\n"))
}
