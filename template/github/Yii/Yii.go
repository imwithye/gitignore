package Yii

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"assets/*",
		"!assets/.gitignore",
		"protected/runtime/*",
		"!protected/runtime/.gitignore",
		"protected/data/*.db",
		"themes/classic/views/",
	}
	template.SetTemplate("GitHub/Yii", strings.Join(ignore, "\n"))
}
