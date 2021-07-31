package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	Name           string `json:"name"`
	Age            int    `json:"age"`
	IgnoreProperty string `json:"-"`
}

func main() {
	u := user{"123", 1, ""}
	fmt.Println(u)

	// convert to json
	ub, err := json.Marshal(u)

	if err != nil {
		log.Fatal(err)
	}

	uj := bytes.NewBuffer(ub)
	fmt.Println(uj)

	// json to struct
	var user user
	js := `{
		"name": "test",
		"age": 1
		}`

	json.Unmarshal([]byte(js), &user)

	fmt.Println(user)
}
