package main

import (
	"errors"
	"strings"
	"unicode"
	"strconv"
)

// Парсер "доверяет" валидатору, поэтому не проверяет явно ошибочные случаи (пустую строку, отсутствием одного или обоих операндов и др.).
// Единственная проверка в парсере -- на наличие оператора, т.к. его наличие -- основа работы парсера.
// На валидатор также опирается реализация утилитарных функций -- предлолагается, что всё, что дошло до парсера будет в ASCII
// В целом, это спорная практика, так как увеличивает связность компонентов. С другой стороны, каждый компонент занимается только своей работой.
// В рамках небольшой учебно-практической работы такой подход считаю допустимым.

func Parser (input string) (int, string, int, bool, error) {
	var leftOperand, rightOperand int
	var operator string
	var isRoman bool

	// Почистим инпут от пробелов, найдем оператор (+,-,*,/) и его индекс во входящем инпуте.
	inputSpaceCleared := removeSpaces(input)
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
	if isDigit(leftStr[0]) {
		leftOperand, _ = strconv.Atoi(leftStr)
		rightOperand, _ = strconv.Atoi(rightStr)
		isRoman = false
		return leftOperand, operator, rightOperand, isRoman, nil
	// В противном случае необходимо преобразовать римские числа к арабским.
	} else {
		leftOperand = romanToInt(leftStr)
		rightOperand = romanToInt(rightStr)
		isRoman = true
		return leftOperand, operator, rightOperand, isRoman, nil
	}
}


func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

//гарантированно чистит инпут от всего, что похоже на пробел по unicode
func removeSpaces(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, s)
}

func romanToInt(s string) int {
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