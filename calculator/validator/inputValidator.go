package validator

import (
	"calculator/constants"
	"errors"
	"regexp"
)

// TODO: минус такой реализации в том, что невозможно понять в чем именно ошибка ввода. Валидатор только говорит, правильно ли введена строка,
// но не показывает причины ошибки, если она есть.

var re = regexp.MustCompile(constants.RomanArabicInputPattern)

func ValidateInput(input string) (bool, error) {
	if re.MatchString(input) {
		return true, nil
	}
	return false, errors.New("введенные данные не могут быть обработаны")
}
