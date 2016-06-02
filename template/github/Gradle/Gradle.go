package Gradle

import "strings"
import "github.com/imwithye/git_ignore/template"

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
	}
	template.SetTemplate("GitHub/Gradle", strings.Join(ignore, "\n"))
}
