package Packer

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"# Cache objects",
		"packer_cache/",
		"",
		"# For built boxes",
		"*.box",
	}
	template.SetTemplate("GitHub/Packer", strings.Join(ignore, "\n"))
}
