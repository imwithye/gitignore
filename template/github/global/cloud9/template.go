package cloud9

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# Cloud9 IDE - http://c9.io",
		".c9revisions",
		".c9",
		"",
	}
	template.Add("github/global/cloud9", strings.Join(ignore, "\n"))
}
