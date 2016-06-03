package haskell

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"dist",
		"dist-*",
		"cabal-dev",
		"*.o",
		"*.hi",
		"*.chi",
		"*.chs.h",
		"*.dyn_o",
		"*.dyn_hi",
		".hpc",
		".hsenv",
		".cabal-sandbox/",
		"cabal.sandbox.config",
		"*.prof",
		"*.aux",
		"*.hp",
		"*.eventlog",
		".stack-work/",
		"cabal.project.local",
		"",
	}
	template.Add("github/haskell", strings.Join(ignore, "\n"))
}
