package TurboGears2

import "strings"
import "github.com/imwithye/git_ignore/template"

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
	}
	template.SetTemplate("GitHub/TurboGears2", strings.Join(ignore, "\n"))
}
