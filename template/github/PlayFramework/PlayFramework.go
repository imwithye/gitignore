package PlayFramework

import "strings"
import "github.com/imwithye/gitignore/template"

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
	}
	template.SetTemplate("GitHub/PlayFramework", strings.Join(ignore, "\n"))
}
