package main

import (
	"fmt"
	"time"
)

func main() {
	go write("hello")
	write("hello2")
}

func write(txt string) {
	for {
		fmt.Println(txt)
		time.Sleep(time.Second)
	}
}
