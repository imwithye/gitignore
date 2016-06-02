package CommonLisp

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"*.FASL",
		"*.fasl",
		"*.lisp-temp",
	}
	template.SetTemplate("GitHub/CommonLisp", strings.Join(ignore, "\n"))
}
