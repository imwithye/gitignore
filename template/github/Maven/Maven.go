package Maven

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"target/",
		"pom.xml.tag",
		"pom.xml.releaseBackup",
		"pom.xml.versionsBackup",
		"pom.xml.next",
		"release.properties",
		"dependency-reduced-pom.xml",
		"buildNumber.properties",
		".mvn/timing.properties",
	}
	template.SetTemplate("GitHub/Maven", strings.Join(ignore, "\n"))
}
