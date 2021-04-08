package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	answer, _ := divide(6, 3)
	fmt.Println("Divide", answer)

	numbers := []float64{10.5, 10, 0.5, 4}
	total := sum(numbers...)
	fmt.Println("sum:", total)
}

func sum(values ...float64) float64 {
	total := 0.0
	for _, value := range values {
		total += value
	}
	return total
}

func divide(p1, p2 float64) (float64, error) {
	if p2 == 0 {
		return math.NaN(), errors.New("cannot divide by zero")
	}
	return p1 / p2, nil
}
