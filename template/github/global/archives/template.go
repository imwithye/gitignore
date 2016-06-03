package archives

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# It's better to unpack these files and commit the raw source because",
		"# git has its own built in compression methods.",
		"*.7z",
		"*.jar",
		"*.rar",
		"*.zip",
		"*.gz",
		"*.bzip",
		"*.bz2",
		"*.xz",
		"*.lzma",
		"*.cab",
		"",
		"#packing-only formats",
		"*.iso",
		"*.tar",
		"",
		"#package management formats",
		"*.dmg",
		"*.xpi",
		"*.gem",
		"*.egg",
		"*.deb",
		"*.rpm",
		"*.msi",
		"*.msm",
		"*.msp",
		"",
	}
	template.Add("github/global/archives", strings.Join(ignore, "\n"))
}
