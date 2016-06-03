package actionscript

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# Build and Release Folders",
		"bin/",
		"bin-debug/",
		"bin-release/",
		"[Oo]bj/ # FlashDevelop obj",
		"[Bb]in/ # FlashDevelop bin",
		"",
		"# Other files and folders",
		".settings/",
		"",
		"# Executables",
		"*.swf",
		"*.air",
		"*.ipa",
		"*.apk",
		"",
		"# Project files, i.e. `.project`, `.actionScriptProperties` and `.flexProperties`",
		"# should NOT be excluded as they contain compiler settings and other important",
		"# information for Eclipse / Flash Builder.",
		"",
	}
	template.Add("github/actionscript", strings.Join(ignore, "\n"))
}
