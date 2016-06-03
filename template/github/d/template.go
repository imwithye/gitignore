package d

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# Compiled Object files",
		"*.o",
		"*.obj",
		"",
		"# Compiled Dynamic libraries",
		"*.so",
		"*.dylib",
		"*.dll",
		"",
		"# Compiled Static libraries",
		"*.a",
		"*.lib",
		"",
		"# Executables",
		"*.exe",
		"",
		"# DUB",
		".dub",
		"docs.json",
		"__dummy.html",
		"",
	}
	template.Add("github/d", strings.Join(ignore, "\n"))
}
