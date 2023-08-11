package utils

import (
	"strings"
	"unicode"
)

// Returns only valid ASCII and UTF8 characters
func BArrayToString(b []byte) string {
	str := string(b)
	onlyASCII := strings.Map(
		func(r rune) rune {
			if r > unicode.MaxASCII {
				return -1
			}
			return r
		},
		str)

	return strings.ToValidUTF8(onlyASCII, "")
}

func RemoveLastQuote(str string) string {
	firstCount := 0
	var outStr string
	for _, c := range str {
		if c == '"' {
			firstCount++
		}
	}
	secondCount := 0
	for _, c := range str {
		if c == '"' {
			secondCount++
			if firstCount == secondCount {
				c = ' '
			}
		}
		outStr += string(c)
	}
	return strings.TrimSpace(outStr)
}
