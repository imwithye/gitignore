package fortran

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# Compiled Object files",
		"*.slo",
		"*.lo",
		"*.o",
		"*.obj",
		"",
		"# Precompiled Headers",
		"*.gch",
		"*.pch",
		"",
		"# Compiled Dynamic libraries",
		"*.so",
		"*.dylib",
		"*.dll",
		"",
		"# Fortran module files",
		"*.mod",
		"*.smod",
		"",
		"# Compiled Static libraries",
		"*.lai",
		"*.la",
		"*.a",
		"*.lib",
		"",
		"# Executables",
		"*.exe",
		"*.out",
		"*.app",
		"",
	}
	template.Add("github/fortran", strings.Join(ignore, "\n"))
}
