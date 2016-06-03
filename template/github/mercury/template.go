package mercury

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"Mercury/",
		"Mercury.modules",
		"*.mh",
		"*.err",
		"*.init",
		"*.dll",
		"*.exe",
		"*.a",
		"*.so",
		"*.dylib",
		"*.beams",
		"*.d",
		"*.c_date",
		"",
	}
	template.Add("github/mercury", strings.Join(ignore, "\n"))
}
