package utils

import (
	"testing"
)

func TestIsDigit(t *testing.T) {
	tests := []struct {
		name     string
		b        byte
		expected bool
	} {
		// Валидные цифры
		{"'0' is digit", '0', true},
        {"'1' is digit", '1', true},

		// Не цифры
		{"'a' is not digit", 'a', false},
        {"'Z' is not digit", 'Z', false},
        {"'-' is not digit", '-', false},
	}

	for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := IsDigit(tt.b)
            if result != tt.expected {
                t.Errorf("IsDigit(%q) = %t, want %t", tt.b, result, tt.expected)
            }
        })
    }
}

func TestRemoveSpaces(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {"Empty string", "", ""},
        {"No spaces", "abc", "abc"},
        {"Only spaces", "   ", ""},
        {"Spaces inside", "a b c", "abc"},
        {"Leading and trailing", "  test  ", "test"},
        {"Multiple spaces", "hello    world", "helloworld"},
        {"Tab", "hello\tworld", "helloworld"},
        {"Newline", "hello\nworld", "helloworld"},
        {"Mixed whitespace", "a \t\n\r b", "ab"},
        {"Emoji with spaces", "👋  🌍  Go", "👋🌍Go"},
        {"Cyrillic", "пр и ве т", "привет"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := RemoveSpaces(tt.input)
            if result != tt.expected {
                t.Errorf("RemoveSpaces(%q) = %q, want %q", tt.input, result, tt.expected)
            }
        })
    }
}

func TestRomanToInt(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected int
    }{
        {"I", "I", 1},
        {"II", "II", 2},
        {"III", "III", 3},
        {"IV", "IV", 4},
        {"V", "V", 5},
        {"VI", "VI", 6},
        {"IX", "IX", 9},
        {"X", "X", 10},
        {"XI", "XI", 11},
        {"XL", "XL", 40},
        {"XC", "XC", 90},
        {"CD", "CD", 400},
        {"CM", "CM", 900},
        {"M", "M", 1000},
        {"MCM", "MCM", 1900},
        {"MCMXC", "MCMXC", 1990},
        {"MMXXIV", "MMXXIV", 2024},
        {"IV+V", "IV+V", 9}, // IV=4, +V=5 → 4+5=9
        {"XIV", "XIV", 14},
        {"XXXIX", "XXXIX", 39},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := RomanToInt(tt.input)
            if result != tt.expected {
                t.Errorf("RomanToInt(%q) = %d, want %d", tt.input, result, tt.expected)
            }
        })
    }
}

func TestIntToRoman(t *testing.T) {
    tests := []struct {
        name         string
        input        int
        expected     string
    }{
        {"1", 1, "I"},
        {"4", 4, "IV"},
        {"5", 5, "V"},
        {"9", 9, "IX"},
        {"10", 10, "X"},
        {"40", 40, "XL"},
        {"50", 50, "L"},
        {"90", 90, "XC"},
        {"100", 100, "C"},
        {"400", 400, "CD"},
        {"500", 500, "D"},
        {"900", 900, "CM"},
        {"1000", 1000, "M"},
        {"1990", 1990, "MCMXC"},
        {"2024", 2024, "MMXXIV"},
        {"3999", 3999, "MMMCMXCIX"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result:= IntToRoman(tt.input)

            if result != tt.expected {
                t.Errorf("IntToRoman(%d) = %q, want %q", tt.input, result, tt.expected)
            }
        })
    }
}