package IPythonNotebook

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"# Temporary data",
		".ipynb_checkpoints/",
	}
	template.SetTemplate("GitHub/Global/IPythonNotebook", strings.Join(ignore, "\n"))
}
