package Vagrant

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		".vagrant/",
	}
	template.SetTemplate("GitHub/Global/Vagrant", strings.Join(ignore, "\n"))
}
