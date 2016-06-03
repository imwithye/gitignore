package jboss

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"jboss/server/all/deploy/project.ext",
		"jboss/server/default/deploy/project.ext",
		"jboss/server/minimal/deploy/project.ext",
		"jboss/server/all/log/*.log",
		"jboss/server/all/tmp/**/*",
		"jboss/server/all/data/**/*",
		"jboss/server/all/work/**/*",
		"jboss/server/default/log/*.log",
		"jboss/server/default/tmp/**/*",
		"jboss/server/default/data/**/*",
		"jboss/server/default/work/**/*",
		"jboss/server/minimal/log/*.log",
		"jboss/server/minimal/tmp/**/*",
		"jboss/server/minimal/data/**/*",
		"jboss/server/minimal/work/**/*",
		"",
		"# deployed package files #",
		"",
		"*.DEPLOYED",
		"",
	}
	template.Add("github/jboss", strings.Join(ignore, "\n"))
}
