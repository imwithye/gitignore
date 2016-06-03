package labview

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

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
		"",
	}
	template.Add("github/labview", strings.Join(ignore, "\n"))
}
