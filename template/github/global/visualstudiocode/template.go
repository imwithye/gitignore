package visualstudiocode

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		".vscode",
		"",
		"",
	}
	template.Add("github/global/visualstudiocode", strings.Join(ignore, "\n"))
}
