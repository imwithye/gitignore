package turbogears2

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"*.py[co]",
		"",
		"# Default development database",
		"devdata.db",
		"",
		"# Default data directory",
		"data/*",
		"",
		"# Packages",
		"*.egg",
		"*.egg-info",
		"dist",
		"build",
		"",
		"# Installer logs",
		"pip-log.txt",
		"",
		"# Unit test / coverage reports",
		".coverage",
		".tox",
		"",
	}
	template.Add("github/turbogears2", strings.Join(ignore, "\n"))
}
