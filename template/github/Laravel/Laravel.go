package Laravel

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"vendor/",
		"node_modules/",
		"",
		"# Laravel 4 specific",
		"bootstrap/compiled.php",
		"app/storage/",
		"",
		"# Laravel 5 & Lumen specific",
		"bootstrap/cache/",
		".env.*.php",
		".env.php",
		".env",
		"",
		"# Rocketeer PHP task runner and deployment package. https://github.com/rocketeers/rocketeer",
		".rocketeer/",
	}
	template.SetTemplate("GitHub/Laravel", strings.Join(ignore, "\n"))
}
