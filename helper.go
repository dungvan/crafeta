package main

import (
	"regexp"
)

var pattern = regexp.MustCompile(`\s+`)

func trimDuplicatedWhiteSpaces(s string) string {
	return pattern.ReplaceAllString(s, " ")
}
