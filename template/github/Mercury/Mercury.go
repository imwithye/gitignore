package Mercury

import "strings"
import "github.com/imwithye/git_ignore/template"

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
	}
	template.SetTemplate("GitHub/Mercury", strings.Join(ignore, "\n"))
}
