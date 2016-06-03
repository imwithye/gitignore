package clojure

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"Leiningen.gitignore",
	}
	template.Add("github/clojure", strings.Join(ignore, "\n"))
}
