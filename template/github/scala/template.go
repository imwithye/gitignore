package scala

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

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
		"",
	}
	template.Add("github/scala", strings.Join(ignore, "\n"))
}
