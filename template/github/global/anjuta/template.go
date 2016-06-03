package anjuta

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# Local configuration folder and symbol database",
		"/.anjuta/",
		"/.anjuta_sym_db.db",
		"",
	}
	template.Add("github/global/anjuta", strings.Join(ignore, "\n"))
}
