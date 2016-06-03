package cfwheels

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

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
		"",
	}
	template.Add("github/cfwheels", strings.Join(ignore, "\n"))
}
