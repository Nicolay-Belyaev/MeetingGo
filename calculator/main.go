package main

// TODO: просмотреть все функции на свежую голову
//       сделать так, что бы калькулятор вместо падения в панику при неправильном вводе или ошибках вычисления перезапускался
// 	     ещё раз подумать, на каком уровне и как именно, и какие именно ошибки обрабатывать


func main() {
	// Получение ввода от пользователя, валидация введенных данных:
	input := GetInput()
	if _, err := validateInput(input); err != nil {
		panic(err)
	}

	// Парсинг ввода пользователя
	leftOperand, operator, rightOperand, isRoman, err := Parser(input)
    if err != nil {
		panic(err)
	}

	result, err := calc(leftOperand, rightOperand, operator)
	if err != nil {
		panic(err)
	}

	OutputResult(result, isRoman)

}