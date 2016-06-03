package jekyll

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"_site/",
		".sass-cache/",
		".jekyll-metadata",
		"",
	}
	template.Add("github/jekyll", strings.Join(ignore, "\n"))
}
