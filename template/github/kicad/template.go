package kicad

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# For PCBs designed using KiCad: http://www.kicad-pcb.org/",
		"",
		"# Temporary files",
		"*.000",
		"*.bak",
		"*.bck",
		"*.kicad_pcb-bak",
		"*~",
		"_autosave-*",
		"*.tmp",
		"",
		"# Netlist files (exported from Eeschema)",
		"*.net",
		"",
		"# Autorouter files (exported from Pcbnew)",
		".dsn",
		"",
		"# Exported BOM files",
		"*.xml",
		"*.csv",
		"",
	}
	template.Add("github/kicad", strings.Join(ignore, "\n"))
}
