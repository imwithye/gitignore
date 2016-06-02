package Archives

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"# It's better to unpack these files and commit the raw source because",
		"# git has its own built in compression methods.",
		"*.7z",
		"*.jar",
		"*.rar",
		"gitignore-master.zip",
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
	}
	template.SetTemplate("GitHub/Global/Archives", strings.Join(ignore, "\n"))
}
