package input

import (
	"bufio"
	"fmt"
	"os"
)

func GetInput() (string, error) {
	var input string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Введите число, затем оператор, затем второе число и нажмите Enter: ")

    if scanner.Scan() {
        input = scanner.Text()
    }

    if err := scanner.Err(); err != nil {
        return "", err
    }

	return input, nil
}
