package textmate

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"*.tmproj",
		"*.tmproject",
		"tmtags",
		"",
	}
	template.Add("github/global/textmate", strings.Join(ignore, "\n"))
}
