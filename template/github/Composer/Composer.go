package Composer

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"composer.phar",
		"/vendor/",
		"",
		"# Commit your application's lock file http://getcomposer.org/doc/01-basic-usage.md#composer-lock-the-lock-file",
		"# You may choose to ignore a library lock file http://getcomposer.org/doc/02-libraries.md#lock-file",
		"# composer.lock",
	}
	template.SetTemplate("GitHub/Composer", strings.Join(ignore, "\n"))
}
