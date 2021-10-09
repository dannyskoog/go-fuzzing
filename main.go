package main

import (
	"fmt"
	"go-fuzzing/internal/pkg/calculation"
)

func main() {
	multiplier := calculation.Multiplier{}

	sum, err := multiplier.Multiply(10, 2)
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
	}

	fmt.Println(sum)
}
