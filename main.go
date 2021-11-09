package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/sahilm/fuzzy"
)

//go:embed templates
var templates embed.FS

func gitignore(args []string) string {
	var ret []string
	for _, arg := range args {
		ret = append(ret, fmt.Sprintf("# Gitignore %s", arg))
		data, err := templates.ReadFile(fmt.Sprintf("templates/%v.gitignore", arg))
		if err != nil {
			continue
		}
		ret = append(ret, string(data))
		ret = append(ret, "\n")
	}
	str := strings.Join(ret, "\n")
	str = strings.ReplaceAll(str, "\r", "\n")
	reg := regexp.MustCompile(`([\n]{3})+`)
	str = reg.ReplaceAllString(str, "\n\n")
	str = strings.ReplaceAll(str, "\n\n", "\n")
	reg = regexp.MustCompile(`\n*$`)
	str = reg.ReplaceAllString(str, "\n")
	return str
}

type IgnoreFile struct {
	Name string
	Path string
}

type IgnoreFiles []IgnoreFile

func (f IgnoreFiles) String(i int) string {
	return f[i].Name
}

func (f IgnoreFiles) Len() int {
	return len(f)
}

func getAllIgnores() IgnoreFiles {
	ignores := IgnoreFiles{}
	fs.WalkDir(templates, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(path, ".gitignore") {
			reg := regexp.MustCompile(`^templates/`)
			path = reg.ReplaceAllString(path, "")
			reg = regexp.MustCompile(`.gitignore$`)
			path = reg.ReplaceAllString(path, "")
			name := filepath.Base(path)
			ignores = append(ignores, IgnoreFile{Name: name, Path: path})
		}
		return nil
	})
	return ignores
}

func fuzzyMatch(arg string, all IgnoreFiles) *IgnoreFile {
	matches := fuzzy.FindFrom(arg, all)
	for _, match := range matches {
		return &all[match.Index]
	}
	return nil
}

func main() {
	args := os.Args[1:]
	matches := []string{}
	for _, arg := range args {
		f := fuzzyMatch(arg, getAllIgnores())
		if f == nil {
			continue
		}
		matches = append(matches, f.Path)
	}
	if len(matches) == 0 {
		fmt.Fprintf(os.Stderr, "No gitignore template found.\n")
		return
	}
	fmt.Printf("Gitignore of [%s] writes to .gitignore file.\n", strings.Join(matches, ", "))
	ioutil.WriteFile(".gitignore", []byte(gitignore(matches)), 0666)
}
