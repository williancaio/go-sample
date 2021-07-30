package main

import (
	"fmt"
	"time"
)

func main() {
	channel1, channel2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Second)
			channel1 <- "channel1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 5)
			channel2 <- "channel2"
		}
	}()

	for {

		select {
		case msg1 := <-channel1:
			fmt.Println(msg1)
		case msg2 := <-channel2:
			fmt.Println(msg2)
		}
	}
}
