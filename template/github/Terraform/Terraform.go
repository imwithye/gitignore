package Terraform

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"# Compiled files",
		"*.tfstate",
		"*.tfstate.backup",
	}
	template.SetTemplate("GitHub/Terraform", strings.Join(ignore, "\n"))
}
