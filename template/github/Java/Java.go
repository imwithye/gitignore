package Java

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"*.class",
		"",
		"# Mobile Tools for Java (J2ME)",
		".mtj.tmp/",
		"",
		"# Package Files #",
		"*.jar",
		"*.war",
		"*.ear",
		"",
		"# virtual machine crash logs, see http://www.java.com/en/download/help/error_hotspot.xml",
		"hs_err_pid*",
	}
	template.SetTemplate("GitHub/Java", strings.Join(ignore, "\n"))
}
