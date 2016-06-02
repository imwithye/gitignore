package GitBook

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"# Node rules:",
		"## Grunt intermediate storage (http://gruntjs.com/creating-plugins#storing-task-files)",
		".grunt",
		"",
		"## Dependency directory",
		"## Commenting this out is preferred by some people, see",
		"## https://docs.npmjs.com/misc/faq#should-i-check-my-node_modules-folder-into-git",
		"node_modules",
		"",
		"# Book build output",
		"_book",
		"",
		"# eBook build output",
		"*.epub",
		"*.mobi",
		"*.pdf",
	}
	template.SetTemplate("GitHub/GitBook", strings.Join(ignore, "\n"))
}
