package Scala

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"*.class",
		"*.log",
		"",
		"# sbt specific",
		".cache",
		".history",
		".lib/",
		"dist/*",
		"target/",
		"lib_managed/",
		"src_managed/",
		"project/boot/",
		"project/plugins/project/",
		"",
		"# Scala-IDE specific",
		".scala_dependencies",
		".worksheet",
	}
	template.SetTemplate("GitHub/Scala", strings.Join(ignore, "\n"))
}
