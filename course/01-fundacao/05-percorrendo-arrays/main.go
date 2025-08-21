package main

import "fmt"

func main() {
	var meuArray [3]int

	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 10

	fmt.Println("Meu ultimo item:", meuArray[len(meuArray)-1])

	fmt.Println("Percorrendo Array")

	for i, v := range meuArray {
		fmt.Printf("O valor do índice %v é %v\n", i, v)
	}
}
