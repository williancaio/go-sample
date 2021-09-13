package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	base64encoded := base64.StdEncoding.EncodeToString([]byte("Hello"))

	fmt.Println("base64 encoded")
	fmt.Println(base64encoded)

	base64decodedbytes, err := base64.StdEncoding.DecodeString(base64encoded)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("base64 decoded bytes")
	fmt.Println(base64decodedbytes)

	base64decodedString := string(base64decodedbytes[:])
	fmt.Println("base64 decoded string")
	fmt.Println(base64decodedString)
}
