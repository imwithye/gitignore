package espresso

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"*.esproj",
		"",
	}
	template.Add("github/global/espresso", strings.Join(ignore, "\n"))
}
