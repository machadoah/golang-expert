package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(sum(1, 5))
	fmt.Println(sub(1, 5))
	fmt.Println(sum_and_verify_even(3, 5))

	valor, isEven, err := sum_positives_and_verify_even(-3, 5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Resultado = %v e o valor é par? %v\n", valor, isEven)
	}
}

func sum(a int, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func sum_and_verify_even(a, b int) (int, bool) {
	sum := a + b
	isEven := (sum % 2) == 0

	return a + b, isEven
}

func sum_positives_and_verify_even(a, b int) (int, bool, error) {
	if a >= 0 && b >= 0 {
		sum := a + b
		isEven := (sum % 2) == 0

		return a + b, isEven, nil
	}
	return 0, false, errors.New("Os valores não são positivos")
}
