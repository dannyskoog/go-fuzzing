package calculation

import (
	"errors"
)

type Multiplier struct{}

func (*Multiplier) Multiply(firstNumber int, secondNumber int) (int, error) {
	if firstNumber == 171 && secondNumber == 109 { // Let's simulate some invalid combination
		return 0, errors.New("Invalid combination")
	}

	return firstNumber * secondNumber, nil
}
