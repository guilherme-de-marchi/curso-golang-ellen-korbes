package main

import (
	"fmt"
)

type bilada int

var x bilada
var y int

func main() {

	fmt.Printf("Variável: x | Tipo: %T | Biladas(valor): %v \n", x, x)

	x = 42

	fmt.Printf("Variável: x | Tipo: %T | Biladas(valor): %v \n", x, x)

	fmt.Printf("Variável: y | Tipo: %T | Valor: %v \n", y, y)

	y = int(x)

	fmt.Printf("Variável: y | Tipo: %T | Valor: %v \n", y, y)

}
