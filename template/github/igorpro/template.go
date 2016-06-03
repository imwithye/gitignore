package igorpro

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# Avoid including Experiment files: they can be created and edited locally to test the ipf files",
		"*.pxp",
		"*.pxt",
		"*.uxp",
		"*.uxt",
		"",
	}
	template.Add("github/igorpro", strings.Join(ignore, "\n"))
}
