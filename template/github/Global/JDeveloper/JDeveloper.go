package JDeveloper

import "strings"
import "github.com/imwithye/gitignore/template"

func init() {
	ignore := []string{
		"# default application storage directory used by the IDE Performance Cache feature",
		".data/",
		"",
		"# used for ADF styles caching",
		"temp/",
		"",
		"# default output directories",
		"classes/",
		"deploy/",
		"javadoc/",
		"",
		"# lock file, a part of Oracle Credential Store Framework",
		"cwallet.sso.lck",
	}
	template.SetTemplate("GitHub/Global/JDeveloper", strings.Join(ignore, "\n"))
}
