package maven

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

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
		"",
	}
	template.Add("github/maven", strings.Join(ignore, "\n"))
}
