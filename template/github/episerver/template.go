package episerver

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"######################",
		"## EPiServer Files",
		"######################",
		"*License.config",
		"",
	}
	template.Add("github/episerver", strings.Join(ignore, "\n"))
}
