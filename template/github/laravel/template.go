package laravel

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

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
		"",
	}
	template.Add("github/laravel", strings.Join(ignore, "\n"))
}
