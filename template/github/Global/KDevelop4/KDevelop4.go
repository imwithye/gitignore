package KDevelop4

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"*.kdev4",
		".kdev4/",
	}
	template.SetTemplate("GitHub/Global/KDevelop4", strings.Join(ignore, "\n"))
}
