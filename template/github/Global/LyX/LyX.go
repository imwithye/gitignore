package LyX

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"# Ignore LyX backup and autosave files",
		"# http://www.lyx.org/",
		"*.lyx~",
		"*.lyx#",
	}
	template.SetTemplate("GitHub/Global/LyX", strings.Join(ignore, "\n"))
}
