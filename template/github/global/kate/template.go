package kate

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# Swap Files #",
		".*.kate-swp",
		".swp.*",
		"",
	}
	template.Add("github/global/kate", strings.Join(ignore, "\n"))
}
