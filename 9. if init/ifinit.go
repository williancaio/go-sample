package main

import "fmt"

func main() {
	num := 1

	if othernum := num; num > 0 {
		othernum++
		if num < othernum {
			fmt.Println("ok")
		} else {
			fmt.Println("nok")
		}
	}
}
