package main

import "fmt"

func main() {
	txt := "main"
	fmt.Println(txt)

	newFunction := closure()
	newFunction()
}

func closure() func() {
	txt := "closure"

	return func() {
		fmt.Println(txt)
	}
}
