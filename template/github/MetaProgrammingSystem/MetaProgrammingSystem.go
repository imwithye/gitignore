package MetaProgrammingSystem

import "strings"
import "github.com/imwithye/git_ignore/template"

func init() {
	ignore := []string{
		"workspace.xml",
		"junitvmwatcher*.properties",
		"build.properties",
		"",
		"# generated java classes and java source files",
		"# manually add any custom artifacts that can't be generated from the models",
		"# http://confluence.jetbrains.com/display/MPSD25/HowTo+--+MPS+and+Git",
		"classes_gen",
		"source_gen",
		"source_gen.caches",
		"",
		"# generated test code and test results",
		"test_gen",
		"test_gen.caches",
		"TEST-*.xml",
		"junit*.properties",
	}
	template.SetTemplate("GitHub/MetaProgrammingSystem", strings.Join(ignore, "\n"))
}
