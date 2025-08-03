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

	// Если в римских, то все, что меньше единицы нельзя показать:
	if result < 1 {
		return errors.New("в римских цифрах нельзя показать результат меньше единицы")
	}
	romanNum := utils.IntToRoman(result)
	fmt.Printf("Когитатор завершил работу. Результат: %s\n", romanNum)
	return nil	
}

