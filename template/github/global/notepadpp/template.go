package notepadpp

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# Notepad++ backups #",
		"*.bak",
		"",
	}
	template.Add("github/global/notepadpp", strings.Join(ignore, "\n"))
}
