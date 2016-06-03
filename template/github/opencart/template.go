package opencart

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		".htaccess",
		"/config.php",
		"admin/config.php",
		"",
		"!index.html",
		"",
		"download/",
		"image/data/",
		"image/cache/",
		"system/cache/",
		"system/logs/",
		"",
		"system/storage/",
		"",
	}
	template.Add("github/opencart", strings.Join(ignore, "\n"))
}
