package Autotools

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"# http://www.gnu.org/software/automake",
		"",
		"Makefile.in",
		"",
		"# http://www.gnu.org/software/autoconf",
		"",
		"/autom4te.cache",
		"/autoscan.log",
		"/autoscan-*.log",
		"/aclocal.m4",
		"/compile",
		"/config.h.in",
		"/configure",
		"/configure.scan",
		"/depcomp",
		"/install-sh",
		"/missing",
		"/stamp-h1",
	}
	template.SetTemplate("GitHub/Autotools", strings.Join(ignore, "\n"))
}
