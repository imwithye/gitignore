package playframework

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# Ignore Play! working directory #",
		"bin/",
		"/db",
		".eclipse",
		"/lib/",
		"/logs/",
		"/modules",
		"/project/target",
		"/target",
		"tmp/",
		"test-result",
		"server.pid",
		"*.eml",
		"/dist/",
		".cache",
		"",
	}
	template.Add("github/playframework", strings.Join(ignore, "\n"))
}
