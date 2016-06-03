package vim

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

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
		"",
	}
	template.Add("github/global/vim", strings.Join(ignore, "\n"))
}
