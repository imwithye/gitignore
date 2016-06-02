package CFWheels

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"# unpacked plugin folders",
		"plugins/**/*",
		"",
		"# files directory where uploads go",
		"files",
		"",
		"# DBMigrate plugin: generated SQL",
		"db/sql",
		"",
		"# AssetBundler plugin: generated bundles",
		"javascripts/bundles",
		"stylesheets/bundles",
	}
	template.SetTemplate("GitHub/CFWheels", strings.Join(ignore, "\n"))
}
