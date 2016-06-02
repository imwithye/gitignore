package Clojure

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"pom.xml",
		"pom.xml.asc",
		"*jar",
		"/lib/",
		"/classes/",
		"/target/",
		"/checkouts/",
		".lein-deps-sum",
		".lein-repl-history",
		".lein-plugins/",
		".lein-failures",
		".nrepl-port",
	}
	template.SetTemplate("GitHub/Clojure", strings.Join(ignore, "\n"))
}
