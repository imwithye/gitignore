package ForceDotCom

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		".project",
		".settings",
		"salesforce.schema",
		"Referenced Packages",
	}
	template.SetTemplate("GitHub/ForceDotCom", strings.Join(ignore, "\n"))
}
