package dreamweaver

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# DW Dreamweaver added files",
		"_notes",
		"_compareTemp",
		"configs/",
		"dwsync.xml",
		"dw_php_codehinting.config",
		"*.mno",
		"",
	}
	template.Add("github/global/dreamweaver", strings.Join(ignore, "\n"))
}
