package EPiServer

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"######################",
		"## EPiServer Files",
		"######################",
		"*License.config",
	}
	template.SetTemplate("GitHub/EPiServer", strings.Join(ignore, "\n"))
}
