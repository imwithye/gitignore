package rhodesrhomobile

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"rholog-*",
		"sim-*",
		"bin/libs",
		"bin/RhoBundle",
		"bin/tmp",
		"bin/target",
		"bin/*.ap_",
		"*.o",
		"*.jar",
		"",
	}
	template.Add("github/rhodesrhomobile", strings.Join(ignore, "\n"))
}
