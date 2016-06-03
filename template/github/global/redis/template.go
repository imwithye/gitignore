package redis

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# Ignore redis binary dump (dump.rdb) files",
		"",
		"*.rdb",
		"",
	}
	template.Add("github/global/redis", strings.Join(ignore, "\n"))
}
