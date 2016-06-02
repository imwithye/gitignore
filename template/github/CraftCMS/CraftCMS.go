package CraftCMS

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"# Craft Storage (cache) [http://buildwithcraft.com/help/craft-storage-gitignore]",
		"/craft/storage/*",
		"!/craft/storage/logo/*",
	}
	template.SetTemplate("GitHub/CraftCMS", strings.Join(ignore, "\n"))
}
