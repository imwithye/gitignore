package Calabash

import "strings"
import "github.com/imwithye/gitignore/template"

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
	}
	template.SetTemplate("GitHub/Global/Calabash", strings.Join(ignore, "\n"))
}
