package stella

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# Atari 2600 (Stella) support for multiple assemblers",
		"# - DASM",
		"# - CC65",
		"",
		"# Assembled binaries and object directories",
		"obj/",
		"a.out",
		"*.bin",
		"*.a26",
		"",
		"# Add in special Atari 7800-based binaries for good measure",
		"*.a78",
		"",
	}
	template.Add("github/stella", strings.Join(ignore, "\n"))
}
