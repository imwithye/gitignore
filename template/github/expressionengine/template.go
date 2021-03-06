package expressionengine

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		".DS_Store",
		"",
		"# Images",
		"images/avatars/",
		"images/captchas/",
		"images/smileys/",
		"images/member_photos/",
		"images/signature_attachments/",
		"images/pm_attachments/",
		"",
		"# For security do not publish the following files",
		"system/expressionengine/config/database.php",
		"system/expressionengine/config/config.php",
		"",
		"# Caches",
		"sized/",
		"thumbs/",
		"_thumbs/",
		"*/expressionengine/cache/*",
		"",
	}
	template.Add("github/expressionengine", strings.Join(ignore, "\n"))
}
