package plone

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"*.pyc",
		"*.pyo",
		"*.tmp*",
		"*.mo",
		"*.egg",
		"*.EGG",
		"*.egg-info",
		"*.EGG-INFO",
		".*.cfg",
		"bin/",
		"build/",
		"develop-eggs/",
		"downloads/",
		"eggs/",
		"fake-eggs/",
		"parts/",
		"dist/",
		"var/",
		"",
	}
	template.Add("github/plone", strings.Join(ignore, "\n"))
}
