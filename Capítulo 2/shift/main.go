package main

import "fmt"

func main() {

	var param_1 int
	fmt.Print("Valor: ")
	fmt.Scan(&param_1)

	fmt.Printf("\nDecimal: %d | Hexadecimal: %x | Binario: %b\n", param_1, param_1, param_1)

	param_2 := param_1 << 1

	fmt.Printf("\nDecimal: %d | Hexadecimal: %x | Binario: %b\n", param_2, param_2, param_2)

}
