package FlexBuilder

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"bin/",
		"bin-debug/",
		"bin-release/",
	}
	template.SetTemplate("GitHub/Global/FlexBuilder", strings.Join(ignore, "\n"))
}
