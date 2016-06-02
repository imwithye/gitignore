package Vagrant

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		".vagrant/",
	}
	template.SetTemplate("GitHub/Global/Vagrant", strings.Join(ignore, "\n"))
}
