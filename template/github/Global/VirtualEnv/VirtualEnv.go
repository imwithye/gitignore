package VirtualEnv

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"# Virtualenv",
		"# http://iamzed.com/2009/05/07/a-primer-on-virtualenv/",
		".Python",
		"[Bb]in",
		"[Ii]nclude",
		"[Ll]ib",
		"[Ll]ib64",
		"[Ll]ocal",
		"[Ss]cripts",
		"pyvenv.cfg",
		".venv",
		"pip-selfcheck.json",
	}
	template.SetTemplate("GitHub/Global/VirtualEnv", strings.Join(ignore, "\n"))
}
