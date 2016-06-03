package julia

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"*.jl.cov",
		"*.jl.*.cov",
		"*.jl.mem",
		"deps/deps.jl",
		"",
	}
	template.Add("github/julia", strings.Join(ignore, "\n"))
}
