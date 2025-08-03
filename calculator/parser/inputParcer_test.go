package parser

import (
	"testing"
)

func TestInputParser(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedLeft  int
		expectedOp    string
		expectedRight int
		expectedRoman bool
		expectError   bool
	}{
		// Арабские числа
		{"Arabic: 2+3", "2+3", 2, "+", 3, false, false},
		{"Arabic: 10-5", "10 - 5", 10, "-", 5, false, false},
		{"Arabic: 4 * 6", "4 * 6", 4, "*", 6, false, false},
		{"Arabic: 8/2", "8/2", 8, "/", 2, false, false},
		{"Arabic: 1*1", "1 * 1", 1, "*", 1, false, false},

		// Римские числа
		{"Roman: I+V", "I+V", 1, "+", 5, true, false},
		{"Roman: X-V", "X - V", 10, "-", 5, true, false},
		{"Roman: II*III", "II * III", 2, "*", 3, true, false},
		{"Roman: VI/II", "VI/II", 6, "/", 2, true, false},
		{"Roman: IV+V", "IV + V", 4, "+", 5, true, false},
		{"Roman: IX*II", "IX * II", 9, "*", 2, true, false},

		// Граничные случаи
		{"Roman: I+I", "I+I", 1, "+", 1, true, false},
		{"Arabic: 10*10", "10*10", 10, "*", 10, false, false},

		// Ошибки: нет оператора
		{"No operator: 123", "123", 0, "", 0, false, true},
		{"No operator: XV", "XV", 0, "", 0, false, true},

		// Пробелы разного типа
		{"With tab", "2\t+\t3", 2, "+", 3, false, false},
		{"With newline", "5\n*\n2", 5, "*", 2, false, false},
		{"Mixed spaces", " 10  /  2 ", 10, "/", 2, false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			left, op, right, roman, err := InputParser(tt.input)

			if tt.expectError {
				if err == nil {
					t.Errorf("Ожидалась ошибка, но её не было")
				}
				return
			}

			if err != nil {
				t.Fatalf("Parser(%q) неожиданная ошибка: %v", tt.input, err)
			}

			if left != tt.expectedLeft {
				t.Errorf("left = %d, want %d", left, tt.expectedLeft)
			}
			if op != tt.expectedOp {
				t.Errorf("op = %q, want %q", op, tt.expectedOp)
			}
			if right != tt.expectedRight {
				t.Errorf("right = %d, want %d", right, tt.expectedRight)
			}
			if roman != tt.expectedRoman {
				t.Errorf("roman = %v, want %v", roman, tt.expectedRoman)
			}
		})
	}
}
