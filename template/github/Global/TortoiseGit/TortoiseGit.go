package TortoiseGit

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"# Project-level settings",
		"/.tgitconfig",
	}
	template.SetTemplate("GitHub/Global/TortoiseGit", strings.Join(ignore, "\n"))
}
