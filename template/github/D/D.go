package D

import "strings"
import "github.com/imwithye/gitignore/template"

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
	}
	template.SetTemplate("GitHub/D", strings.Join(ignore, "\n"))
}
