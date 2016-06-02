package LemonStand

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"boot.php",
		"index.php",
		"install.php",
		"/config/*",
		"!/config/config.php",
		"/controllers/*",
		"/init/*",
		"/logs/*",
		"/phproad/*",
		"/temp/*",
		"/uploaded/*",
		"/installer_files/*",
		"/modules/backend/*",
		"/modules/blog/*",
		"/modules/cms/*",
		"/modules/core/*",
		"/modules/session/*",
		"/modules/shop/*",
		"/modules/system/*",
		"/modules/users/*",
		"# add content_*.php if you don't want erase client changes to content",
	}
	template.SetTemplate("GitHub/LemonStand", strings.Join(ignore, "\n"))
}
