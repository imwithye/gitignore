package MicrosoftOffice

import "strings"
import "github.com/imwithye/gitignore/template"

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
	}
	template.SetTemplate("GitHub/Global/MicrosoftOffice", strings.Join(ignore, "\n"))
}
