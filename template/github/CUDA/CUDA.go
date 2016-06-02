package CUDA

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"*.i",
		"*.ii",
		"*.gpu",
		"*.ptx",
		"*.cubin",
		"*.fatbin",
	}
	template.SetTemplate("GitHub/CUDA", strings.Join(ignore, "\n"))
}
