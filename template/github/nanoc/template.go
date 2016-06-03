package nanoc

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

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
		"",
	}
	template.Add("github/nanoc", strings.Join(ignore, "\n"))
}
