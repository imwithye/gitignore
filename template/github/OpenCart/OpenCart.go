package OpenCart

import "strings"
import "github.com/imwithye/gitignore/template"

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
	}
	template.SetTemplate("GitHub/OpenCart", strings.Join(ignore, "\n"))
}
