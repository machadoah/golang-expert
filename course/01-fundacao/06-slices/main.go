package main

import "fmt"

func main() {
	s := []int{2, 4, 6, 8, 10}

	fmt.Printf("len=%v cap=%v slice=%v\n", len(s), cap(s), s)

	// slice vazio
	fmt.Printf("len=%v cap=%v slice=%v\n", len(s[:0]), cap(s[:0]), s[:0])

	fmt.Printf("len=%v cap=%v slice=%v\n", len(s[:4]), cap(s[:4]), s[:4])

	// diminuiu em 2 a capacidade
	fmt.Printf("len=%v cap=%v slice=%v\n", len(s[2:]), cap(s[2:]), s[2:])

	s = append(s, 12)

	// capacidade vai de 5 para 10, inicial * 2
	fmt.Printf("len=%v cap=%v slice=%v\n", len(s), cap(s), s)
}
