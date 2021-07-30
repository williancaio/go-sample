package main

import "fmt"

func main() {
	var num1 int = 1
	var pointer *int = &num1
	*pointer = 3

	fmt.Println(num1)
}
