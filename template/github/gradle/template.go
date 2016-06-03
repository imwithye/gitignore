package gradle

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		".gradle",
		"build/",
		"",
		"# Ignore Gradle GUI config",
		"gradle-app.setting",
		"",
		"# Avoid ignoring Gradle wrapper jar file (.jar files are usually ignored)",
		"!gradle-wrapper.jar",
		"",
		"# Cache of project",
		".gradletasknamecache",
		"",
		"# # Work around https://youtrack.jetbrains.com/issue/IDEA-116898",
		"# gradle/wrapper/gradle-wrapper.properties",
		"",
	}
	template.Add("github/gradle", strings.Join(ignore, "\n"))
}
