package template

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

var NotExist = errors.New("ignore template does not exist")
var template map[string]string

type Template struct {
	Ignore string
	Path   string
}

func Add(key, val string) {
	if template == nil {
		template = map[string]string{}
	}
	template[strings.ToLower(key)] = val
}

func Get(key string) (Template, error) {
	if template == nil {
		return Template{notExist(key), ""}, notExistError(key)
	}
	if path := Matchs(key); path == "" {
		return Template{notExist(key), ""}, notExistError(key)
	} else if val, ok := template[path]; !ok {
		return Template{notExist(key), ""}, notExistError(key)
	} else {
		return Template{val, path}, nil
	}
}

func Matchs(k string) string {
	for key, _ := range template {
		if strings.ToLower(k) == filepath.Base(key) {
			return key
		}
	}
	return ""
}

func notExist(key string) string {
	return fmt.Sprintf("# gitignore template for %s does not exist", key)
}

func notExistError(key string) error {
	return errors.New(fmt.Sprintf("gitignore template for %s does not exist", key))
}
