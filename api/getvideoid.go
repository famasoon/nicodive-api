package api

import "regexp"

func GetVideoID(link string) string {
	r := regexp.MustCompile(`([a-z][a-z])[0-9]+$`)
	return r.FindAllStringSubmatch(link, -1)[0][0]
}
