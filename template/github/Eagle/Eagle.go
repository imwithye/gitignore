package Eagle

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"# Ignore list for Eagle, a PCB layout tool",
		"",
		"# Backup files",
		"*.s#?",
		"*.b#?",
		"*.l#?",
		"",
		"# Eagle project file",
		"# It contains a serial number and references to the file structure",
		"# on your computer.",
		"# comment the following line if you want to have your project file included.",
		"eagle.epf",
		"",
		"# Autorouter files",
		"*.pro",
		"*.job",
		"",
		"# CAM files",
		"*.$$$",
		"*.cmp",
		"*.ly2",
		"*.l15",
		"*.sol",
		"*.plc",
		"*.stc",
		"*.sts",
		"*.crc",
		"*.crs",
		"",
		"*.dri",
		"*.drl",
		"*.gpi",
		"*.pls",
		"",
		"*.drd",
		"*.drd.*",
		"",
		"*.info",
		"",
		"*.eps",
		"",
		"# file locks introduced since 7.x",
		"*.lck",
		"",
	}
	template.SetTemplate("GitHub/Eagle", strings.Join(ignore, "\n"))
}
