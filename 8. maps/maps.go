package main

import "fmt"

func main() {
	user := map[string]string{
		"name":      "your name",
		"last_name": "your last name",
	}

	fmt.Println(user)

	user2 := map[string]map[string]string{
		"name": {
			"first": "haha",
			"last":  "hehe",
		},
	}

	fmt.Println(user2)

	delete(user, "last_name")
	fmt.Println(user)
}
