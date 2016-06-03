package vagrant

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		".vagrant/",
		"",
	}
	template.Add("github/global/vagrant", strings.Join(ignore, "\n"))
}
