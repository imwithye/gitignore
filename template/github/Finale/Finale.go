package Finale

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"*.bak",
		"*.db",
		"*.avi",
		"*.pdf",
		"*.ps",
		"*.mid",
		"*.midi",
		"*.mp3",
		"*.aif",
		"*.wav",
		"# Some versions of Finale have a bug and randomly save extra copies of",
		"# the music source as \"<Filename> copy.mus\"",
		"*copy.mus",
	}
	template.SetTemplate("GitHub/Finale", strings.Join(ignore, "\n"))
}
