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
