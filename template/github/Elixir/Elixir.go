package Elixir

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"/_build",
		"/cover",
		"/deps",
		"erl_crash.dump",
		"*.ez",
	}
	template.SetTemplate("GitHub/Elixir", strings.Join(ignore, "\n"))
}
