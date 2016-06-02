package Erlang

import "strings"
import "github.com/imwithye/git_ignore/template"

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
	}
	template.SetTemplate("GitHub/Erlang", strings.Join(ignore, "\n"))
}
