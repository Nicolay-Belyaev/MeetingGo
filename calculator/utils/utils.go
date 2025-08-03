package utils

import (
	"strings"
	"unicode"
)

func IsDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

// гарантированно чистит инпут от всего, что похоже на пробел по unicode
func RemoveSpaces(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, s)
}

func RomanToInt(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && romanMap[s[i]] < romanMap[s[i+1]] {
			result -= romanMap[s[i]]
		} else {
			result += romanMap[s[i]]
		}
	}

	return result
}

func IntToRoman(n int) string {
    // Список значений и соответствующих римских символов,
    // включая вычитательные пары (IV, IX, XL, XC, CD, CM)
    values := []int{
        1000, 900, 500, 400,
        100, 90, 50, 40,
        10, 9, 5, 4,
        1,
    }
    numerals := []string{
        "M", "CM", "D", "CD",
        "C", "XC", "L", "XL",
        "X", "IX", "V", "IV",
        "I",
    }

    var result strings.Builder
    for i := 0; i < len(values); i++ {
        count := n / values[i]
        for j := 0; j < count; j++ {
            result.WriteString(numerals[i])
        }
        n -= values[i] * count
    }

    return result.String()
}