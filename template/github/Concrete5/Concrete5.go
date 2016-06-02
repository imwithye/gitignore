package Concrete5

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"config/site.php",
		"files/cache/*",
		"files/tmp/*",
		".htaccess",
	}
	template.SetTemplate("GitHub/Concrete5", strings.Join(ignore, "\n"))
}
