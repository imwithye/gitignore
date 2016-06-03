package sketchup

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"*.skb",
		"",
	}
	template.Add("github/sketchup", strings.Join(ignore, "\n"))
}
