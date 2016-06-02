package Xcode

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"# Xcode",
		"#",
		"# gitignore contributors: remember to update Global/Xcode.gitignore, Objective-C.gitignore & Swift.gitignore",
		"",
		"## Build generated",
		"build/",
		"DerivedData/",
		"",
		"## Various settings",
		"*.pbxuser",
		"!default.pbxuser",
		"*.mode1v3",
		"!default.mode1v3",
		"*.mode2v3",
		"!default.mode2v3",
		"*.perspectivev3",
		"!default.perspectivev3",
		"xcuserdata/",
		"",
		"## Other",
		"*.moved-aside",
		"*.xccheckout",
		"*.xcscmblueprint",
	}
	template.SetTemplate("GitHub/Global/Xcode", strings.Join(ignore, "\n"))
}
