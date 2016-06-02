package Smalltalk

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"# changes file",
		"*.changes",
		"",
		"# system image",
		"*.image",
		"",
		"# Pharo Smalltalk Debug log file",
		"PharoDebug.log",
		"",
		"# Squeak Smalltalk Debug log file",
		"SqueakDebug.log",
		"",
		"# Monticello package cache",
		"/package-cache",
		"",
		"# Metacello-github cache",
		"/github-cache",
		"github-*.zip",
	}
	template.SetTemplate("GitHub/Smalltalk", strings.Join(ignore, "\n"))
}
