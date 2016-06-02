package ZendFramework

import "strings"
import "github.com/imwithye/git_ignore/template"

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
		"",
		"# Legacy ZF1",
		"demos/",
		"extras/documentation",
	}
	template.SetTemplate("GitHub/ZendFramework", strings.Join(ignore, "\n"))
}
