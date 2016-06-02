package Anjuta

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"# Local configuration folder and symbol database",
		"/.anjuta/",
		"/.anjuta_sym_db.db",
	}
	template.SetTemplate("GitHub/Global/Anjuta", strings.Join(ignore, "\n"))
}
