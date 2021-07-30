package main

import "fmt"

func main() {
	var variable1 string = "one string"
	fmt.Println(variable1)

	variable2 := "inference string"
	fmt.Println(variable2)

	var (
		variable3 string
		variable4 string
	)

	variable3 = "var 3"
	variable4 = "var 4"
	fmt.Println(variable3)
	fmt.Println(variable4)

	variable5, variable6 := 1, 2

	fmt.Println(variable6, variable5)
}
