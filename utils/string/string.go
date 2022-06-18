package string

import (
	"encoding/json"
	"log"
	"regexp"
	"strings"
)

func PrintJson(title string, v interface{}) {
	m, _ := json.Marshal(v)
	log.Println(title, string(m))
}

func FilterName(username string) string {
	return regexp.MustCompile("[^\u4e00-\u9fa5|[A-Za-z]|\\d]").ReplaceAllString(username, "")
}

func Join(s string, sep string) string {
	var arr []string
	for _, i := range s {
		arr = append(arr, string(i))
	}
	return strings.Join(arr, sep)
}
