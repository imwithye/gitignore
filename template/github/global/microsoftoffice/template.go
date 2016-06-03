package microsoftoffice

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"*.tmp",
		"",
		"# Word temporary",
		"~$*.doc*",
		"",
		"# Excel temporary",
		"~$*.xls*",
		"",
		"# Excel Backup File",
		"*.xlk",
		"",
		"# PowerPoint temporary",
		"~$*.ppt*",
		"",
		"# Visio autosave temporary files",
		"*.~vsdx",
		"",
	}
	template.Add("github/global/microsoftoffice", strings.Join(ignore, "\n"))
}
