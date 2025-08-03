package parser

import (
	"errors"
	"strconv"
	"calculator/utils"
)

// Парсер "доверяет" валидатору, поэтому не проверяет явно ошибочные случаи (пустую строку, отсутствием одного или обоих операндов и др.).
// Единственная проверка в парсере -- на наличие оператора, т.к. его наличие -- основа работы парсера.
// На валидатор также опирается реализация утилитарных функций -- предлолагается, что всё, что дошло до парсера будет в ASCII
// Это спорная практика, так как увеличивает связность компонентов. С другой стороны, каждый компонент занимается только своей работой.
// В рамках небольшой учебно-практической работы такой подход считаю допустимым.

func InputParser(input string) (int, string, int, bool, error) {
	var leftOperand, rightOperand int
	var operator string
	var isRoman bool

	// Почистим инпут от пробелов, найдем оператор (+,-,*,/) и его индекс во входящем инпуте.
	inputSpaceCleared := utils.RemoveSpaces(input)
	operatorIndex := -1
	for i, r := range inputSpaceCleared {
		// TODO: подумать над регуляркой!
		if r == '+' || r == '-' || r == '*' || r == '/' {
			operatorIndex = i
			operator = string(r)
			break
		}
	}

	// Если валидатор как-то пропустил некорректную строку, в которой нет оператора -- выйдем с ошибокой.
	if operatorIndex == -1 {
		return 0, "", 0, isRoman, errors.New("ошибка парсера: не найден математический оператор")
	}

	// Так как у нас нет пробелов в очищенной строке, все, что слева от оператора будет одним операндом, все, что справа -- другим.
	var leftStr string = inputSpaceCleared[:operatorIndex]
	var rightStr string = inputSpaceCleared[operatorIndex+1:]

	// Если первый символ левого операнда арабская цифра, то достаточно привести тип операндов к int.
	if utils.IsDigit(leftStr[0]) {
		leftOperand, _ := strconv.Atoi(leftStr)
		rightOperand, _ := strconv.Atoi(rightStr)
		isRoman = false
		return leftOperand, operator, rightOperand, isRoman, nil
		// В противном случае необходимо преобразовать римские числа к арабским.
	} else {
		leftOperand = utils.RomanToInt(leftStr)
		rightOperand = utils.RomanToInt(rightStr)
		isRoman = true
		return leftOperand, operator, rightOperand, isRoman, nil
	}
}
