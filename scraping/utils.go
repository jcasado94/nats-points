package scraping

import "strings"

func cleanTag(tag string) string {
	return strings.ToLower(tag)
}
