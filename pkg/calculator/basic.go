package calculator

import (
	"errors"
	"github.com/joaopapereira/versioned-module/v3/pkg/converter"
)

func AddString(num1 string, num2 string) (int, error) {
	a, err := converter.StrToInt(num1)
	if err != nil {
		return 0, errors.Join(errors.New("conversion first number"), err)
	}
	b, err := converter.StrToInt(num2)
	if err != nil {
		return 0, errors.Join(errors.New("conversion second number"), err)
	}
	return Add(a, b), nil
}

func Add(a, b int) int {
	return a + b
}

func SubString(num1 string, num2 string) (int, error) {
	a, err := converter.StrToInt(num1)
	if err != nil {
		return 0, errors.Join(errors.New("conversion first number"), err)
	}
	b, err := converter.StrToInt(num2)
	if err != nil {
		return 0, errors.Join(errors.New("conversion second number"), err)
	}
	return Sub(a, b), nil
}
func Sub(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}

	return a / b, nil
}
