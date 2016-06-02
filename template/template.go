package template

import (
	"strings"
)

var template map[string]string

func SetTemplate(key, val string) {
	if template == nil {
		template = map[string]string{}
	}
	template[strings.ToLower(key)] = val
}

func GetTemplate(key string) string {
	if template == nil {
		return ""
	}
	if val, ok := template[strings.ToLower(key)]; ok {
		return val
	} else {
		return ""
	}
}
