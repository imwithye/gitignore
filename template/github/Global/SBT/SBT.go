package SBT

import "strings"
import "github.com/imwithye/gitignore/template"

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
	}
	template.SetTemplate("GitHub/Global/SBT", strings.Join(ignore, "\n"))
}
