package IGORPro

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"# Avoid including Experiment files: they can be created and edited locally to test the ipf files",
		"*.pxp",
		"*.pxt",
		"*.uxp",
		"*.uxt",
	}
	template.SetTemplate("GitHub/IGORPro", strings.Join(ignore, "\n"))
}
