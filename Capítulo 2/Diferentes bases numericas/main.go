package main

import "fmt"

func main() {

	var value int
	fmt.Print("Valor: ")
	fmt.Scan(&value)

	fmt.Printf("\nDecimal:     %d\n", value)
	fmt.Printf("Hexadecimal: %x\n", value)
	fmt.Printf("Binário:     %b\n", value)
	fmt.Printf("Octagonal:   %o\n", value)

}
