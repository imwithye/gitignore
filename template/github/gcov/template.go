package gcov

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# gcc coverage testing tool files",
		"",
		"*.gcno",
		"*.gcda",
		"*.gcov",
		"",
	}
	template.Add("github/gcov", strings.Join(ignore, "\n"))
}
