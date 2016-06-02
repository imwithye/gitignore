package LibreOffice

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"# LibreOffice locks",
		".~lock.*#",
	}
	template.SetTemplate("GitHub/Global/LibreOffice", strings.Join(ignore, "\n"))
}
