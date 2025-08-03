package main

import (
	"errors"
	"fmt"
	"strings"
)

func OutputResult(result int, isRoman bool) error {
	// Если мы в арабских цифрах:
	if !isRoman {
		fmt.Printf("Когитатор завершил работу. Результат: %d\n", result)
		return nil
	}

	// Если в римских, то условие сложнее:
	if result < 1 {
		return errors.New("когитатор завершил работу. В римских цифрах нельзя показать результат меньше единицы")
	}
	romanNum, err := intToRoman(result)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Когитатор завершил работу. Результат: %s\n", romanNum)
	return nil	
}

func intToRoman(n int) (string, error) {
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

    return result.String(), nil
}