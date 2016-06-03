package lyx

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# Ignore LyX backup and autosave files",
		"# http://www.lyx.org/",
		"*.lyx~",
		"*.lyx#",
		"",
	}
	template.Add("github/global/lyx", strings.Join(ignore, "\n"))
}
