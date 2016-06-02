package main

import (
	"archive/zip"
	"fmt"
	"github.com/imwithye/logor"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const PackageFolder = "template"
const PackageName = "template"

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

func generateCode(path string) {
	log := logor.GetLogor()
	relPath, err := filepath.Rel("gitignore-master", path)
	if err != nil {
		log.Fatal("   ", err)
		return
	}
	extension := filepath.Ext(relPath)
	packagePath := filepath.Join(PackageFolder, relPath[0:len(relPath)-len(extension)])
	packagePath = strings.ToLower(packagePath)
	packagePath = strings.Replace(packagePath, "+", "p", -1)
	packagePath = strings.Replace(packagePath, "-", "_", -1)
	fmt.Println(packagePath)
}

func main() {
	log := logor.GetLogor()
	var err error
	log.Info("-->", "Downloading Github Gitignore Template")
	err = downloadFile("gitignore-master.zip", "https://codeload.github.com/github/gitignore/zip/master")
	if err != nil {
		log.Fatal("-->", err)
	}
	log.Info("-->", "Extracting Github Gitignore Template")
	err = unzipFile("gitignore-master.zip", ".")
	if err != nil {
		log.Fatal("-->", err)
	}
	log.Info("-->", "Generating Go Codes & Packages")
	err = filepath.Walk("gitignore-master", visitFile)
	if err != nil {
		log.Fatal("-->", err)
	}
	log.Info("-->", "Done")
}
