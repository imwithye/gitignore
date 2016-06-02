package TextMate

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"*.tmproj",
		"*.tmproject",
		"tmtags",
	}
	template.SetTemplate("GitHub/Global/TextMate", strings.Join(ignore, "\n"))
}
