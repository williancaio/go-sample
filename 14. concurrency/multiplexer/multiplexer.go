package main

import (
	"fmt"
	"time"
)

func main() {
	channel := multiplexer(write("test1"), write("test2"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func multiplexer(channel1, channel2 <-chan string) <-chan string {
	channelOutput := make(chan string)
	go func() {
		for {
			select {
			case msg := <-channel1:
				channelOutput <- msg
			case msg := <-channel2:
				channelOutput <- msg
			}
		}
	}()
	return channelOutput
}

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- text
			time.Sleep(20000)
		}

	}()

	return channel
}
