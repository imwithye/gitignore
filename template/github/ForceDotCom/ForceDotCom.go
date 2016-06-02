package ForceDotCom

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		".project",
		".settings",
		"salesforce.schema",
		"Referenced Packages",
	}
	template.SetTemplate("GitHub/ForceDotCom", strings.Join(ignore, "\n"))
}
