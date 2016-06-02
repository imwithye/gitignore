package LabVIEW

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"# Libraries",
		"*.lvlibp",
		"*.llb",
		"",
		"# Shared objects (inc. Windows DLLs)",
		"*.dll",
		"*.so",
		"*.so.*",
		"*.dylib",
		"",
		"# Executables",
		"*.exe",
		"",
		"# Metadata",
		"*.aliases",
		"*.lvlps",
	}
	template.SetTemplate("GitHub/LabVIEW", strings.Join(ignore, "\n"))
}
