package clojure

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

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
		"",
	}
	template.Add("github/clojure", strings.Join(ignore, "\n"))
}
