package main

import "fmt"

var x int
var y string
var z bool

func main() {

	fmt.Println("Valores zero de cada tipo: ")
	fmt.Printf("\tVariável: X | Valor: %v | Tipo: %T\n", x, x)
	fmt.Printf("\tVariável: Y | Valor: \"%v\" | Tipo: %T\n", y, y)
	fmt.Printf("\tVariável: Z | Valor: %v | Tipo: %T\n", z, z)

}
