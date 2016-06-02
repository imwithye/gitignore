package ArchLinuxPackages

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"*.tar",
		"*.tar.*",
		"*.jar",
		"*.exe",
		"*.msi",
		"gitignore-master.zip",
		"*.tgz",
		"*.log",
		"*.log.*",
		"*.sig",
		"",
		"pkg/",
		"src/",
	}
	template.SetTemplate("GitHub/ArchLinuxPackages", strings.Join(ignore, "\n"))
}
