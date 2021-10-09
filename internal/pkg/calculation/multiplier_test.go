//go:build go1.18
// +build go1.18

package calculation

import (
	"testing"
)

func FuzzMultiplierMultiply(f *testing.F) {
	f.Add(70, 35)
	f.Add(120, 60)

	f.Fuzz(func(t *testing.T, firstNumber int, secondNumber int) {
		multiplier := &Multiplier{}

		_, err := multiplier.Multiply(firstNumber, secondNumber)

		if err != nil {
			t.Fatalf("Failed to multiply %d with %d: %v", firstNumber, secondNumber, err)
		}
	})
}
