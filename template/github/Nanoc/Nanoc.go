package Nanoc

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"# For projects using nanoc (http://nanoc.ws/)",
		"",
		"# Default location for output, needs to match output_dir's value found in config.yaml",
		"output/",
		"",
		"# Temporary file directory",
		"tmp/",
		"",
		"# Crash Log",
		"crash.log",
	}
	template.SetTemplate("GitHub/Nanoc", strings.Join(ignore, "\n"))
}
