package ocaml

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"*.annot",
		"*.cmo",
		"*.cma",
		"*.cmi",
		"*.a",
		"*.o",
		"*.cmx",
		"*.cmxs",
		"*.cmxa",
		"",
		"# ocamlbuild working directory",
		"_build/",
		"",
		"# ocamlbuild targets",
		"*.byte",
		"*.native",
		"",
		"# oasis generated files",
		"setup.data",
		"setup.log",
		"",
	}
	template.Add("github/ocaml", strings.Join(ignore, "\n"))
}
