package WordPress

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"*.log",
		"wp-config.php",
		"wp-content/advanced-cache.php",
		"wp-content/backup-db/",
		"wp-content/backups/",
		"wp-content/blogs.dir/",
		"wp-content/cache/",
		"wp-content/upgrade/",
		"wp-content/uploads/",
		"wp-content/wp-cache-config.php",
		"wp-content/plugins/hello.php",
		"",
		"/.htaccess",
		"/license.txt",
		"/readme.html",
		"/sitemap.xml",
		"/sitemap.xml.gz",
		"",
	}
	template.SetTemplate("GitHub/WordPress", strings.Join(ignore, "\n"))
}
