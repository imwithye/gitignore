package ModelSim

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"# ignore ModelSim generated files and directories (temp files and so on)",
		"[_@]*",
		"",
		"# ignore compilation output of ModelSim",
		"*.mti",
		"*.dat",
		"*.dbs",
		"*.psm",
		"*.bak",
		"*.cmp",
		"*.jpg",
		"*.html",
		"*.bsf",
		"",
		"# ignore simulation output of ModelSim",
		"wlf*",
		"*.wlf",
		"*.vstf",
		"*.ucdb",
		"cov*/",
		"transcript*",
		"sc_dpiheader.h",
		"vsim.dbg",
	}
	template.SetTemplate("GitHub/Global/ModelSim", strings.Join(ignore, "\n"))
}
