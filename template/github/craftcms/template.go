package craftcms

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# Craft Storage (cache) [http://buildwithcraft.com/help/craft-storage-gitignore]",
		"/craft/storage/*",
		"!/craft/storage/logo/*",
	}
	template.Add("github/craftcms", strings.Join(ignore, "\n"))
}
