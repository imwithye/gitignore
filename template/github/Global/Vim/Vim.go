package Vim

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"# swap",
		"[._]*.s[a-w][a-z]",
		"[._]s[a-w][a-z]",
		"# session",
		"Session.vim",
		"# temporary",
		".netrwhist",
		"*~",
		"# auto-generated tag files",
		"tags",
	}
	template.SetTemplate("GitHub/Global/Vim", strings.Join(ignore, "\n"))
}
