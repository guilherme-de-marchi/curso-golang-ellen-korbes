package main

import "fmt"

func main() {

	x := 42
	y := "James Bond"
	z := true

	fmt.Println("\nVários prints juntos: ")
	fmt.Println(fmt.Sprintf("\t%v %v %v", x, y, z))

	fmt.Println("\nPrints separados agora: ")
	fmt.Println(fmt.Sprintf("\tValor de x: %v", x))
	fmt.Println(fmt.Sprintf("\tValor de y: %v", y))
	fmt.Println(fmt.Sprintf("\tValor de z: %v\n", z))

}
