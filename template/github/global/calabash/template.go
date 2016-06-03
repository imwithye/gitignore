package calabash

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# Calabash / Cucumber",
		"rerun/",
		"reports/",
		"screenshots/",
		"screenshot*.png",
		"test-servers/",
		"",
		"# bundler",
		".bundle",
		"vendor",
		"",
	}
	template.Add("github/global/calabash", strings.Join(ignore, "\n"))
}
