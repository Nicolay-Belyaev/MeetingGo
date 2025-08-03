package output

import (
	"errors"
	"fmt"
	"calculator/utils"
)

func OutputResult(result int, isRoman bool) error {
	// Если мы в арабских цифрах:
	if !isRoman {
		fmt.Printf("Когитатор завершил работу. Результат: %d\n", result)
		return nil
	}

	// Если в римских, то условие сложнее:
	if result < 1 {
		return errors.New("в римских цифрах нельзя показать результат меньше единицы")
	}
	romanNum, err := utils.IntToRoman(result)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Когитатор завершил работу. Результат: %s\n", romanNum)
	return nil	
}

