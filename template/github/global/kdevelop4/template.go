package kdevelop4

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"*.kdev4",
		".kdev4/",
		"",
	}
	template.Add("github/global/kdevelop4", strings.Join(ignore, "\n"))
}
