package MonoDevelop

import "strings"
import "github.com/imwithye/git_ignore/template"

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
	}
	template.SetTemplate("GitHub/Global/MonoDevelop", strings.Join(ignore, "\n"))
}
