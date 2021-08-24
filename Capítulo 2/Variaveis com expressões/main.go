package main

import "fmt"

func main() {

	var param_1 int
	fmt.Print("Primeiro parâmetro: ")
	fmt.Scan(&param_1)

	var param_2 int
	fmt.Print("Segundo parâmetro: ")
	fmt.Scan(&param_2)

	equal := param_1 == param_2
	not_equal := param_1 != param_2
	lower_or_equal := param_1 <= param_2
	higher_or_equal := param_1 >= param_2
	lower := param_1 < param_2
	higher := param_1 > param_2

	fmt.Println("")

	fmt.Printf("Equal: %v\n", equal)
	fmt.Printf("Not Equal: %v\n", not_equal)
	fmt.Printf("Lower or Equal: %v\n", lower_or_equal)
	fmt.Printf("Higher or Equal: %v\n", higher_or_equal)
	fmt.Printf("Lower: %v\n", lower)
	fmt.Printf("Higher: %v\n", higher)

}
