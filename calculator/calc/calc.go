package calc

import "errors"

func Calculate(a int, b int, op string) (int, error) {
    switch op {
		case "+":
			return sum(a, b)
		case "-":
			return subtract(a, b)
		case "*":
			return multiply(a, b)
		case "/":
			return Divide(a, b)
		default:
			return 0, errors.New("неизвестный оператор")
	}
}

func sum(a, b int) (int, error) {
    return a + b, nil
}

func subtract(a, b int) (int, error) {
    return a - b, nil
}

func multiply(a, b int) (int, error) {
    return a * b, nil
}

func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("нельзя делить на ноль")
    }
    return a / b, nil
}