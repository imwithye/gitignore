package Linux

import "strings"
import "github.com/imwithye/gitignore/template"

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
	}
	template.SetTemplate("GitHub/Global/Linux", strings.Join(ignore, "\n"))
}
