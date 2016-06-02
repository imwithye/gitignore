package ChefCookbook

import "strings"
import "github.com/imwithye/gitignore/template"

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
	}
	template.SetTemplate("GitHub/ChefCookbook", strings.Join(ignore, "\n"))
}
