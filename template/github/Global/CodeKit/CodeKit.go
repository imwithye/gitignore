package CodeKit

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"# General CodeKit files to ignore",
		"config.codekit",
		"/min",
	}
	template.SetTemplate("GitHub/Global/CodeKit", strings.Join(ignore, "\n"))
}
