package packer

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# Cache objects",
		"packer_cache/",
		"",
		"# For built boxes",
		"*.box",
		"",
	}
	template.Add("github/packer", strings.Join(ignore, "\n"))
}
