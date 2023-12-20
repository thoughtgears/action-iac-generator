package main

import (
	"strings"
	"unicode"
)

func sanitiseOutputFile(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) || r == '-' {
			return '_'
		}
		return r
	}, s)
}
