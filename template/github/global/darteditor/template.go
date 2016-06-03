package darteditor

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		".project",
		".buildlog",
		"",
	}
	template.Add("github/global/darteditor", strings.Join(ignore, "\n"))
}
