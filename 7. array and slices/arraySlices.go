package main

import "fmt"

func main() {
	var arr [5]string
	fmt.Println(arr)

	arr1 := [5]string{"1", "2"}
	fmt.Println(arr1)

	arr2 := [...]int{1, 2, 34, 5, 6}
	fmt.Println(arr2)

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slice)

	slice = append(slice, 3)

}
