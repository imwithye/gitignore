package CakePHP

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"# CakePHP 3",
		"",
		"/vendor/*",
		"/config/app.php",
		"",
		"/tmp/cache/models/*",
		"!/tmp/cache/models/empty",
		"/tmp/cache/persistent/*",
		"!/tmp/cache/persistent/empty",
		"/tmp/cache/views/*",
		"!/tmp/cache/views/empty",
		"/tmp/sessions/*",
		"!/tmp/sessions/empty",
		"/tmp/tests/*",
		"!/tmp/tests/empty",
		"",
		"/logs/*",
		"!/logs/empty",
		"",
		"# CakePHP 2",
		"",
		"/app/tmp/*",
		"/app/Config/core.php",
		"/app/Config/database.php",
		"/vendors/*",
	}
	template.SetTemplate("GitHub/CakePHP", strings.Join(ignore, "\n"))
}
