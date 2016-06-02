package Gcov

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"# gcc coverage testing tool files",
		"",
		"*.gcno",
		"*.gcda",
		"*.gcov",
	}
	template.SetTemplate("GitHub/Gcov", strings.Join(ignore, "\n"))
}
