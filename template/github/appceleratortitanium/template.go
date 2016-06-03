package appceleratortitanium

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# Build folder and log file",
		"build/",
		"build.log",
		"",
	}
	template.Add("github/appceleratortitanium", strings.Join(ignore, "\n"))
}
