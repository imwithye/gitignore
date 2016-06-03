package qooxdoo

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"cache",
		"cache-downloads",
		"inspector",
		"api",
		"source/inspector.html",
		"",
	}
	template.Add("github/qooxdoo", strings.Join(ignore, "\n"))
}
