package libreoffice

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# LibreOffice locks",
		".~lock.*#",
		"",
	}
	template.Add("github/global/libreoffice", strings.Join(ignore, "\n"))
}
