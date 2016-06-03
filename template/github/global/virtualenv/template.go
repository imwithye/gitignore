package virtualenv

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

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
		"",
	}
	template.Add("github/global/virtualenv", strings.Join(ignore, "\n"))
}
