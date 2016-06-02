package NotepadPP

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"# Notepad++ backups #",
		"*.bak",
	}
	template.SetTemplate("GitHub/Global/NotepadPP", strings.Join(ignore, "\n"))
}
