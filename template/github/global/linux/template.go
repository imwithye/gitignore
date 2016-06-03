package linux

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"*~",
		"",
		"# temporary files which can be created if a process still has a handle open of a deleted file",
		".fuse_hidden*",
		"",
		"# KDE directory preferences",
		".directory",
		"",
		"# Linux trash folder which might appear on any partition or disk",
		".Trash-*",
		"",
	}
	template.Add("github/global/linux", strings.Join(ignore, "\n"))
}