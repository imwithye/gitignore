package modelsim

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

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
		"",
	}
	template.Add("github/global/modelsim", strings.Join(ignore, "\n"))
}
