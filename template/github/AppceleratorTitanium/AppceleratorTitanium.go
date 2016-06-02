package AppceleratorTitanium

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"# Build folder and log file",
		"build/",
		"build.log",
	}
	template.SetTemplate("GitHub/AppceleratorTitanium", strings.Join(ignore, "\n"))
}
