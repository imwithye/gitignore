package java

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

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
		"",
	}
	template.Add("github/java", strings.Join(ignore, "\n"))
}
