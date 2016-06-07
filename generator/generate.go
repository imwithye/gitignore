package main

import (
	"archive/zip"
	"fmt"
	"github.com/imwithye/logor"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const PackageFolder = "template"
const GoTemplate = `package %s

import (
	"strings"
	"github.com/imwithye/gitignore/template"
)

func init() {
	ignore := []string{
%s
	}
	template.Add("%s", strings.Join(ignore, "\n"))
}
`
const ReadmeTemplate = "Git Template %s\n===\n\nUse `git ignore add %s` to add this ignore template.\n\n```\n%s```\n"
const AllTemplate = `package %s

import (
%s
)
`

var imports = []string{}

func downloadFile(filepath string, url string) (err error) {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func unzipFile(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()
		path := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			f, err := os.OpenFile(
				path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()
			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func visitFile(path string, f os.FileInfo, err error) error {
	if err == nil && filepath.Ext(path) == ".gitignore" {
		generateCode(path)
	}
	return nil
}

func symbolic(path string) string {
	log := logor.GetLogor()
	baseName := filepath.Base(path)
	dir := filepath.Dir(path)
	switch baseName {
	case "Fortran.gitignore":
		log.Info("   ", "rewrite path", path, "to", filepath.Join(dir, "C++.gitignore"))
		return filepath.Join(dir, "C++.gitignore")
	case "Clojure.gitignore":
		log.Info("   ", "rewrite path", path, "to", filepath.Join(dir, "Leiningen.gitignore"))
		return filepath.Join(dir, "Leiningen.gitignore")
	default:
		return path
	}
}

func generateCode(path string) {
	log := logor.GetLogor()
	content, err := ioutil.ReadFile(symbolic(path))
	if err != nil {
		log.Error("   ", err)
		return
	}
	ignores := []string{}
	for _, line := range strings.Split(string(content), "\n") {
		line = strings.Replace(line, "\\", "\\\\", -1)
		line = strings.Replace(line, "\"", "\\\"", -1)
		ignores = append(ignores, fmt.Sprintf("\t\t\"%s\",", strings.Trim(line, " ")))
	}
	if ignores[len(ignores)-1] != "" {
		ignores = append(ignores, "")
	}
	relPath, err := filepath.Rel("gitignore-master", path)
	if err != nil {
		log.Error("   ", err)
		return
	}
	extension := filepath.Ext(relPath)
	packagePath := filepath.Join(PackageFolder, "github", relPath[0:len(relPath)-len(extension)])
	packagePath = strings.ToLower(packagePath)
	packagePath = strings.Replace(packagePath, "+", "p", -1)
	packagePath = strings.Replace(packagePath, "-", "_", -1)
	packageName := filepath.Base(packagePath)
	if packageName == "go" {
		packageName = "golang"
	}
	err = os.MkdirAll(packagePath, 0755)
	if err != nil {
		log.Error("   ", err)
		return
	}
	goFilePath := filepath.Join(packagePath, "template.go")
	gocode := fmt.Sprintf(
		GoTemplate, packageName,
		strings.Join(ignores, "\n"),
		strings.ToLower(strings.Replace(filepath.Join("github", relPath[0:len(relPath)-len(extension)]), "\\", "/", -1)))
	ioutil.WriteFile(goFilePath, []byte(gocode), 0644)
	readmeFilePath := filepath.Join(packagePath, "README.md")
	readme := fmt.Sprintf(ReadmeTemplate, relPath[0:len(relPath)-len(extension)], packageName, string(content))
	ioutil.WriteFile(readmeFilePath, []byte(readme), 0644)
	importPath := strings.Replace(fmt.Sprintf("_ \"github.com/imwithye/gitignore/%s\"", packagePath), "\\", "/", -1)
	imports = append(imports, importPath)
}

func main() {
	log := logor.GetLogor()
	var err error
	log.Info("-->", "Downloading Github Gitignore Template")
	err = os.RemoveAll("gitignore-master.zip")
	if err != nil {
		log.Fatal("-->", err)
	}
	err = downloadFile("gitignore-master.zip", "https://codeload.github.com/github/gitignore/zip/master")
	if err != nil {
		log.Fatal("-->", err)
	}
	log.Info("-->", "Extracting Github Gitignore Template")
	err = os.RemoveAll("gitignore-master")
	if err != nil {
		log.Fatal("-->", err)
	}
	err = unzipFile("gitignore-master.zip", ".")
	if err != nil {
		log.Fatal("-->", err)
	}
	log.Info("-->", "Generating Go Codes & Packages")
	err = os.RemoveAll(filepath.Join(PackageFolder, "github"))
	if err != nil {
		log.Fatal("-->", err)
	}
	err = os.RemoveAll(filepath.Join(PackageFolder, "all"))
	if err != nil {
		log.Fatal("-->", err)
	}
	err = filepath.Walk("gitignore-master", visitFile)
	if err != nil {
		log.Fatal("-->", err)
	}
	allPath := filepath.Join(PackageFolder, "all")
	err = os.MkdirAll(allPath, 0755)
	if err != nil {
		log.Fatal("-->", err)
	}
	allPath = filepath.Join(PackageFolder, "all", "all.go")
	err = ioutil.WriteFile(
		allPath,
		[]byte(fmt.Sprintf(AllTemplate, "all", strings.Join(imports, "\n"))),
		0644)
	if err != nil {
		log.Fatal("-->", err)
	}
	log.Info("-->", "Done")
}
