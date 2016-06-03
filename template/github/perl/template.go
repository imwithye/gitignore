package perl

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"/blib/",
		"/.build/",
		"_build/",
		"cover_db/",
		"inc/",
		"Build",
		"!Build/",
		"Build.bat",
		".last_cover_stats",
		"/Makefile",
		"/Makefile.old",
		"/MANIFEST.bak",
		"/META.yml",
		"/META.json",
		"/MYMETA.*",
		"nytprof.out",
		"/pm_to_blib",
		"*.o",
		"*.bs",
		"/_eumm/",
		"",
	}
	template.Add("github/perl", strings.Join(ignore, "\n"))
}
