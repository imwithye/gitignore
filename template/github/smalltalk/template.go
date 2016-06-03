package smalltalk

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

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
		"",
	}
	template.Add("github/smalltalk", strings.Join(ignore, "\n"))
}
