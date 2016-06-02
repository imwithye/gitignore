package EPiServer

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"######################",
		"## EPiServer Files",
		"######################",
		"*License.config",
	}
	template.SetTemplate("GitHub/EPiServer", strings.Join(ignore, "\n"))
}
