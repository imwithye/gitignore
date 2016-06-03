package windows

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"# Windows image file caches",
		"Thumbs.db",
		"ehthumbs.db",
		"",
		"# Folder config file",
		"Desktop.ini",
		"",
		"# Recycle Bin used on file shares",
		"$RECYCLE.BIN/",
		"",
		"# Windows Installer files",
		"*.cab",
		"*.msi",
		"*.msm",
		"*.msp",
		"",
		"# Windows shortcuts",
		"*.lnk",
		"",
	}
	template.Add("github/global/windows", strings.Join(ignore, "\n"))
}
