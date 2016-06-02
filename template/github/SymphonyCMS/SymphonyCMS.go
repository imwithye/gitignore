package SymphonyCMS

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"manifest/cache/",
		"manifest/logs/",
		"manifest/tmp/",
		"symphony/",
		"workspace/uploads/",
		"install-log.txt",
	}
	template.SetTemplate("GitHub/SymphonyCMS", strings.Join(ignore, "\n"))
}
