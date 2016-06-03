package elixir

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"/_build",
		"/cover",
		"/deps",
		"erl_crash.dump",
		"*.ez",
		"",
	}
	template.Add("github/elixir", strings.Join(ignore, "\n"))
}
