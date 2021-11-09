package args

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	Outfile   = kingpin.Flag("output", "The output file. Default is .gitignore and use - for stdout.").Default(".gitignore").String()
	Templates = kingpin.Arg("templates", "Gitignore templates").Required().Strings()
)

func Parse() {
	kingpin.Parse()
}
