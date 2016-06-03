package xojo

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

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
		"",
	}
	template.Add("github/xojo", strings.Join(ignore, "\n"))
}
