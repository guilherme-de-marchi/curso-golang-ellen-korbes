package main

import "fmt"

type bilada int

var x bilada

func main() {

	fmt.Printf("Variável: x | Tipo: %T | Biladas(valor): %v \n", x, x)

	x = 42

	fmt.Printf("Variável: x | Tipo: %T | Biladas(valor): %v \n", x, x)

}
