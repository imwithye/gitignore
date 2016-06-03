package concrete5

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"config/site.php",
		"files/cache/*",
		"files/tmp/*",
		".htaccess",
		"",
	}
	template.Add("github/concrete5", strings.Join(ignore, "\n"))
}
