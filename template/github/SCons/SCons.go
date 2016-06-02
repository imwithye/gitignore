package SCons

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"# for projects that use SCons for building: http://http://www.scons.org/",
		".sconsign.dblite",
	}
	template.SetTemplate("GitHub/SCons", strings.Join(ignore, "\n"))
}
