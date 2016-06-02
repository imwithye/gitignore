package Cpp

import "strings"
import "github.com/imwithye/git_ignore/template"

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
	}
	template.SetTemplate("GitHub/Cpp", strings.Join(ignore, "\n"))
}
