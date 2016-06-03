package cuda

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"*.i",
		"*.ii",
		"*.gpu",
		"*.ptx",
		"*.cubin",
		"*.fatbin",
		"",
	}
	template.Add("github/cuda", strings.Join(ignore, "\n"))
}
