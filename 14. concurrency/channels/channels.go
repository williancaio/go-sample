package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	go write("hello", channel)

	for {
		msg, isOpen := <-channel

		fmt.Println(msg)

		if !isOpen {
			break
		}
	}
	channel = make(chan string)
	go write("hello2", channel)
	for mens := range channel {
		fmt.Println(mens)
	}
}

func write(txt string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- txt
		time.Sleep(time.Second)
	}
	close(channel)
}
