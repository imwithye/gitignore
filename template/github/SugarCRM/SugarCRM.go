package SugarCRM

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"## SugarCRM",
		"# Ignore custom .htaccess stuff.",
		"/.htaccess",
		"# Ignore the cache directory completely.",
		"# This will break the current behaviour. Which was often leading to",
		"# the misuse of the repository as backup replacement.",
		"# For development the cache directory can be safely ignored and",
		"# therefore it is ignored.",
		"/cache/",
		"# Ignore some files and directories from the custom directory.",
		"/custom/history/",
		"/custom/modulebuilder/",
		"/custom/working/",
		"/custom/modules/*/Ext/",
		"/custom/application/Ext/",
		"# Custom configuration should also be ignored.",
		"/config.php",
		"/config_override.php",
		"# The silent upgrade scripts aren't needed.",
		"/silentUpgrade*.php",
		"# Logs files can safely be ignored.",
		"*.log",
		"# Ignore the new upload directories.",
		"/upload/",
		"/upload_backup/",
	}
	template.SetTemplate("GitHub/SugarCRM", strings.Join(ignore, "\n"))
}
