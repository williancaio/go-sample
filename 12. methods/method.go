package main

import "fmt"

type user struct {
	name string
	age  int
}

func (u user) save() {
	fmt.Println("ok")
}

func main() {
	user1 := user{"teste", 1}

	user1.save()
}
