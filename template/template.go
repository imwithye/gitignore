package template

import (
	"errors"
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
		return "", NotExist
	}
	if path := Matchs(key); path == "" {
		return "", NotExist
	} else if val, ok := template[path]; !ok {
		return "", NotExist
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
