package main

import "fmt"

type RA int

var (
	nome     string = "Antonio"
	registro RA     = 2727272727
)

func main() {
	fmt.Printf("O tipo de nome é %T\n", nome)
	fmt.Printf("O valor de nome é %v\n", nome)

	fmt.Printf("O tipo de registro é %T\n", registro)
}
