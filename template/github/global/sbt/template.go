package sbt

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# Simple Build Tool",
		"# http://www.scala-sbt.org/release/docs/Getting-Started/Directories.html#configuring-version-control",
		"",
		"target/",
		"lib_managed/",
		"src_managed/",
		"project/boot/",
		".history",
		".cache",
		"",
	}
	template.Add("github/global/sbt", strings.Join(ignore, "\n"))
}
