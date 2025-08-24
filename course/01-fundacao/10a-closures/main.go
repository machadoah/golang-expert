package main

import "fmt"

func main() {
	fmt.Println("*️⃣ Contador 01 criado ...")
	c1 := counter()

	fmt.Println("⚙️ Executando contador 01:")
	fmt.Println(c1()) // 1
	fmt.Println(c1()) // 2
	fmt.Println(c1()) // 3

	fmt.Println("*️⃣ Contador 02 criado ...")
	c2 := counter()

	fmt.Println("⚙️ Executando contador 02:")
	fmt.Println(c2()) // 1
	fmt.Println(c2()) // 2

	fmt.Println("⚙️ Executando o contador 01 novamente:")
	fmt.Println(c1()) // 4
}

func counter() func() int {
	i := 0

	fmt.Println("*️⃣ Variável 'i' inicializada com 0.")

	return func() int {
		i++
		return i
	}
}
