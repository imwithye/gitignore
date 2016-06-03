package cmake

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"CMakeCache.txt",
		"CMakeFiles",
		"CMakeScripts",
		"Makefile",
		"cmake_install.cmake",
		"install_manifest.txt",
		"CTestTestfile.cmake",
		"",
	}
	template.Add("github/cmake", strings.Join(ignore, "\n"))
}
