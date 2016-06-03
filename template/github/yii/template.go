package yii

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"assets/*",
		"!assets/.gitignore",
		"protected/runtime/*",
		"!protected/runtime/.gitignore",
		"protected/data/*.db",
		"themes/classic/views/",
		"",
	}
	template.Add("github/yii", strings.Join(ignore, "\n"))
}
