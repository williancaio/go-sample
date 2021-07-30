package main

import (
	"fmt"
	"time"
)

func main() {
	channel := write("hehe")

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- text
			time.Sleep(time.Second)
		}

	}()

	return channel
}
