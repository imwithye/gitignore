package C

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"# Object files",
		"*.o",
		"*.ko",
		"*.obj",
		"*.elf",
		"",
		"# Precompiled Headers",
		"*.gch",
		"*.pch",
		"",
		"# Libraries",
		"*.lib",
		"*.a",
		"*.la",
		"*.lo",
		"",
		"# Shared objects (inc. Windows DLLs)",
		"*.dll",
		"*.so",
		"*.so.*",
		"*.dylib",
		"",
		"# Executables",
		"*.exe",
		"*.out",
		"*.app",
		"*.i*86",
		"*.x86_64",
		"*.hex",
		"",
		"# Debug files",
		"*.dSYM/",
		"*.su",
	}
	template.SetTemplate("GitHub/C", strings.Join(ignore, "\n"))
}
