package main

import (
	"calculator/calc"
	"calculator/input"
	"calculator/parser"
	"calculator/validator"
	"calculator/output"
	"fmt"
)

func calcLoop() {
	for {
		userInput, inpurErr := input.GetInput()
		if inpurErr != nil {
			fmt.Printf("Ошибка на этапе ввода данных: %v.\n", inpurErr)
			fmt.Println("Попробуйте ещё раз.")
			continue
		}
		if userInput == "exit" || userInput == "quit" {
			fmt.Println("До свидания!")
			break
		}
		if _, validateErr := validator.ValidateInput(userInput); validateErr != nil {
			// Плохое решение с точки зрения UX -- непонятно, в чем именно ошибка. Для детализации надо переписать валидатор.
			fmt.Printf("Ошибка на этапе проверки данных: %v.\n", validateErr)
			fmt.Println("Попробуйте ещё раз.")
			continue
		}
		leftOperand, operator, rightOperand, isRoman, parseErr := parser.InputParser(userInput)
		if parseErr != nil {
			fmt.Printf("Ошибка на этапе чтения данных: %v.\n", parseErr)
			fmt.Println("Попробуйте ещё раз.")
			calcLoop()
		}
		result, calcErr := calc.Calculate(leftOperand, rightOperand, operator)
		if calcErr != nil {
			fmt.Printf("Ошибка на этапе вычислений: %v.\n", calcErr)
			fmt.Println("Попробуйте ещё раз.")
			continue
		}
		outErr := output.OutputResult(result, isRoman)
		if outErr != nil {
			fmt.Printf("Ошибка на этапе вывода результата: %v.\n", outErr)
			fmt.Println("Попробуйте ещё раз.")
			continue
		}
	}
}

func main() {
	calcLoop()
}
