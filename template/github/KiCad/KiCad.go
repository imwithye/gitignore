package KiCad

import "strings"
import "github.com/imwithye/git_ignore/template"

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
	}
	template.SetTemplate("GitHub/KiCad", strings.Join(ignore, "\n"))
}
