package ada

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# Object file",
		"*.o",
		"",
		"# Ada Library Information",
		"*.ali",
		"",
	}
	template.Add("github/ada", strings.Join(ignore, "\n"))
}
