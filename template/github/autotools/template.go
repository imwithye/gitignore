package autotools

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

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
		"",
	}
	template.Add("github/autotools", strings.Join(ignore, "\n"))
}
