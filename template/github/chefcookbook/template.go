package chefcookbook

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		".vagrant",
		"/cookbooks",
		"",
		"# Bundler",
		"bin/*",
		".bundle/*",
		"",
		".kitchen/",
		".kitchen.local.yml",
		"",
	}
	template.Add("github/chefcookbook", strings.Join(ignore, "\n"))
}
