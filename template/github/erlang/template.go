package erlang

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		".eunit",
		"deps",
		"*.o",
		"*.beam",
		"*.plt",
		"erl_crash.dump",
		"ebin",
		"rel/example_project",
		".concrete/DEV_MODE",
		".rebar",
		"",
	}
	template.Add("github/erlang", strings.Join(ignore, "\n"))
}
