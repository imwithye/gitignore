package Redis

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"# Ignore redis binary dump (dump.rdb) files",
		"",
		"*.rdb",
	}
	template.SetTemplate("GitHub/Global/Redis", strings.Join(ignore, "\n"))
}
