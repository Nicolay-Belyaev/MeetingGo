package main

import (
	"calculator/calc"
	"calculator/input"
	"calculator/parser"
	"calculator/validator"
)

// TODO: просмотреть все функции на свежую голову
//       сделать так, что бы калькулятор вместо падения в панику при неправильном вводе или ошибках вычисления перезапускался
// 	     ещё раз подумать, на каком уровне и как именно, и какие именно ошибки обрабатывать

func main() {
	// Получение ввода от пользователя, валидация введенных данных:
	input := input.GetInput()
	if _, err := validator.ValidateInput(input); err != nil {
		panic(err)
	}

	// Парсинг ввода пользователя
	leftOperand, operator, rightOperand, isRoman, err := parser.InputParser(input)
	if err != nil {
		panic(err)
	}

	result, err := calc.Calculate(leftOperand, rightOperand, operator)
	if err != nil {
		panic(err)
	}

	OutputResult(result, isRoman)

}
