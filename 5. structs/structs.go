package main

import "fmt"

type user struct {
	name    string
	age     int
	address address
}

type address struct {
	street string
}

type person struct {
	user
	cpf string
}

func main() {
	var user1 user

	user1.age = 1
	user1.name = "123"

	address := address{"street"}
	user2 := user{"1", 1, address}

	fmt.Println(user2)

	user3 := user{age: 1}

	fmt.Println(user3)

	var person1 person
	person2 := person{user2, "1223"}
	fmt.Println(person1)
	fmt.Println(person2)

}
