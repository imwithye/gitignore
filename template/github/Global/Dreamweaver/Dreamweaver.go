package Dreamweaver

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"# DW Dreamweaver added files",
		"_notes",
		"_compareTemp",
		"configs/",
		"dwsync.xml",
		"dw_php_codehinting.config",
		"*.mno",
	}
	template.SetTemplate("GitHub/Global/Dreamweaver", strings.Join(ignore, "\n"))
}
