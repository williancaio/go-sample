package main

import "fmt"

func main() {
	arr := []string{"val1", "val2", "val3"}

	for i, value := range arr {
		fmt.Println(i, value)
	}

	for _, value := range arr {
		fmt.Println(value)
	}
}
