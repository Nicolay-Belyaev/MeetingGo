package validator

import (
    "testing"
)

func TestValidateInput(t *testing.T) {
    tests := []struct {
        name         string
        input        string
        expectValid  bool
        expectError  bool
    }{
        //  Валидные: римские
        {"Valid: I+V", "I+V", true, false},
        {"Valid: I + V", "I + V", true, false},
        {"Valid: X-V", "X-V", true, false},
        {"Valid: IV*II", "IV*II", true, false},
        {"Valid: IX / III", "IX / III", true, false},

        //  Валидные: арабские
        {"Valid: 1+2", "1+2", true, false},
        {"Valid: 10 * 5", "10 * 5", true, false},
        {"Valid: 2/1", "2/1", true, false},
        {"Valid: 7   -   1", "7   -   1", true, false},

        //  Невалидные: смешанные системы
        {"Invalid: I+2", "I+2", false, true},
        {"Invalid: 5-V", "5-V", false, true},

        //  Невалидные: за пределами диапазона
        {"Invalid: XI+I", "XI+I", false, true},
        {"Invalid: 0+5", "0+5", false, true},
        {"Invalid: 11+1", "11+1", false, true},

        //  Невалидные: некорректные римские
        {"Invalid: IIII+V", "IIII+V", false, true},
        {"Invalid: IIIX+I", "IIIX+I", false, true},

        //  Невалидные: нет оператора
        {"Invalid: no op: XV", "XV", false, true},
        {"Invalid: no op: 12", "12", false, true},

        //  Невалидные: пустое или неполное
        {"Invalid: empty", "", false, true},
        {"Invalid: only left", "X+", false, true},
        {"Invalid: only right", "+V", false, true},
        {"Invalid: spaces only", "   ", false, true},

        // Невалидные: посторонние символы
        {"Invalid: letters", "a+b", false, true},
        {"Invalid: special", "!+@", false, true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            valid, err := ValidateInput(tt.input)

            if valid != tt.expectValid {
                t.Errorf("validateInput(%q): ожидалось валидность %t, получено %t", tt.input, tt.expectValid, valid)
            }

            if tt.expectError {
                if err == nil {
                    t.Errorf("validateInput(%q): ожидалась ошибка, но её не было", tt.input)
                }
            } else {
                if err != nil {
                    t.Errorf("validateInput(%q): неожиданная ошибка: %v", tt.input, err)
                }
            }
        })
    }
}