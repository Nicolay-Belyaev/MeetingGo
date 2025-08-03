package main

import (
	"strings"
	"unicode"
)

// BEGIN (write your solution here)
func LatinLetters(s string) string {
  	resStr := strings.Builder{}
	for _, ch := range s {
		if unicode.Is(unicode.Latin, ch) {
			resStr.WriteRune(ch)
		}
	}
	return resStr.String()
}

// END
