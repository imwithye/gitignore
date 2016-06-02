package template

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

var NotExist = errors.New("ignore template does not exist")
var template map[string]string

func SetTemplate(key, val string) {
	if template == nil {
		template = map[string]string{}
	}
	template[strings.ToLower(key)] = val
}

func GetTemplate(key string) (string, error) {
	if template == nil {
		return notExist(key), notExistError(key)
	}
	if path := Matchs(key); path == "" {
		return notExist(key), notExistError(key)
	} else if val, ok := template[path]; !ok {
		return notExist(key), notExistError(key)
	} else {
		return val, nil
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
	return fmt.Sprintf("# ignore template for %s does not exist", key)
}

func notExistError(key string) error {
	return errors.New(fmt.Sprintf("ignore template for %s does not exist", key))
}
