package calc

import (
	"testing"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		name           string
		operandA       int
		operandB       int
		operator       string
		expectedResult int
		expectedError  bool
	}{
		// Базовые случаи
		{"Sum: 5+3", 5, 3, "+", 8, false},
		{"Subtract: 4-1", 4, 1, "-", 3, false},
		{"Multiply: 2*6", 2, 6, "*", 12, false},
		{"Divide: 4/2", 4, 2, "/", 2, false},

		//Некорректный оператор
		{"Wrong operator", 1, 8, "^", 0, true},

		//Деление на ноль
		{"Divide by zero", 2, 0, "/", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Calculate(tt.operandA, tt.operandB, tt.operator)

			if tt.expectedError {
				if err == nil {
					t.Errorf("Ожидалась ошибка, но её не было")
				}
				if tt.operator == "/" && tt.operandB == 0 {
        			if err.Error() != "нельзя делить на ноль" {
            			t.Errorf("Неверное сообщение об ошибке: %v", err)
        			}
    			}
				return
			}

		if err != nil {
				t.Fatalf("Calc(%d, %d, %s) неожиданная ошибка: %v", tt.operandA, tt.operandB, tt.operator, err)
			}

		if result != tt.expectedResult {
			t.Errorf("result = %d, want %d", result, tt.expectedResult)
		}
		})
	}
}