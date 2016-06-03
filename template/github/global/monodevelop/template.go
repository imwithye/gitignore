package monodevelop

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"#User Specific",
		"*.userprefs",
		"*.usertasks",
		"",
		"#Mono Project Files",
		"*.pidb",
		"*.resources",
		"test-results/",
		"",
	}
	template.Add("github/global/monodevelop", strings.Join(ignore, "\n"))
}
