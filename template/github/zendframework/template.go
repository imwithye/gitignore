package zendframework

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# Composer files",
		"composer.phar",
		"vendor/",
		"",
		"# Local configs",
		"config/autoload/*.local.php",
		"",
		"# Binary gettext files",
		"*.mo",
		"",
		"# Data",
		"data/logs/",
		"data/cache/",
		"data/sessions/",
		"data/tmp/",
		"temp/",
		"",
		"#Doctrine 2",
		"data/DoctrineORMModule/Proxy/",
		"data/DoctrineORMModule/cache/",
		"",
		"# Legacy ZF1",
		"demos/",
		"extras/documentation",
		"",
	}
	template.Add("github/zendframework", strings.Join(ignore, "\n"))
}
