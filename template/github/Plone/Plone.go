package Plone

import "strings"
import "github.com/imwithye/git_ignore/template"

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
	}
	template.SetTemplate("GitHub/Plone", strings.Join(ignore, "\n"))
}
