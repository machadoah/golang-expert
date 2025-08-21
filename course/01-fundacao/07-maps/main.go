package main

import "fmt"

func main() {
	salarios := map[string]int{"Antonio": 1000, "Isadora": 5000, "Ilda": 10000}

	// remove a ilda
	delete(salarios, "Ilda")

	// cria a maria
	salarios["Maria"] = 7000

	for name, salario := range salarios {
		fmt.Printf("O salário de %v é de %v\n", name, salario)
	}

	idades := make(map[string]int)

	idades["abreu"] = 65
	idades["bernardo"] = 23

	// blank identifier (_)
	for _, idade := range idades {
		fmt.Printf("As idades das pessoas são %v\n", idade)
	}
}
