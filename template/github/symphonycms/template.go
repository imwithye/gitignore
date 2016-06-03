package symphonycms

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"manifest/cache/",
		"manifest/logs/",
		"manifest/tmp/",
		"symphony/",
		"workspace/uploads/",
		"install-log.txt",
		"",
	}
	template.Add("github/symphonycms", strings.Join(ignore, "\n"))
}
