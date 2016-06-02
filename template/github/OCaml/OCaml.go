package OCaml

import "strings"
import "github.com/imwithye/git_ignore/template"

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
	}
	template.SetTemplate("GitHub/OCaml", strings.Join(ignore, "\n"))
}
