package codeigniter

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"*/config/development",
		"*/logs/log-*.php",
		"!*/logs/index.html",
		"*/cache/*",
		"!*/cache/index.html",
		"!*/cache/.htaccess",
		"",
	}
	template.Add("github/codeigniter", strings.Join(ignore, "\n"))
}
