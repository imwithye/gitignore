package scons

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# for projects that use SCons for building: http://http://www.scons.org/",
		".sconsign.dblite",
		"",
	}
	template.Add("github/scons", strings.Join(ignore, "\n"))
}
