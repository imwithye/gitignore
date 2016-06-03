package archlinuxpackages

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"*.tar",
		"*.tar.*",
		"*.jar",
		"*.exe",
		"*.msi",
		"*.zip",
		"*.tgz",
		"*.log",
		"*.log.*",
		"*.sig",
		"",
		"pkg/",
		"src/",
		"",
	}
	template.Add("github/archlinuxpackages", strings.Join(ignore, "\n"))
}
