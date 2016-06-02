package TortoiseGit

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"# Project-level settings",
		"/.tgitconfig",
	}
	template.SetTemplate("GitHub/Global/TortoiseGit", strings.Join(ignore, "\n"))
}
