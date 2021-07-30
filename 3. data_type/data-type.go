package main

import (
	"errors"
	"fmt"
)

func main() {
	var num int = 123
	fmt.Println(num)
	var unum uint = 123
	fmt.Println(unum)

	//alias int32
	var num3 rune = 123
	fmt.Println(num3)

	var num4 byte = 123
	fmt.Println(num4)

	var numberFloat float32 = 123.3
	fmt.Println(numberFloat)

	var numberFloat1 float64 = 123.3
	fmt.Println(numberFloat1)

	var string string = "123"
	fmt.Println(string)

	// default value is false
	var boolean1 bool = true
	fmt.Println(boolean1)

	var error error
	fmt.Println(error)

	error = errors.New("Internal error")
	fmt.Println(error)

}
