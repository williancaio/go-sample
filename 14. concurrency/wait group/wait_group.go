package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		write("hello")
		waitGroup.Done()
	}()

	go func() {
		write("hello2")
		waitGroup.Done()
	}()

	waitGroup.Wait()
	fmt.Println("finish")
}

func write(txt string) {
	for i := 0; i < 5; i++ {
		fmt.Println(txt)
		time.Sleep(time.Second)
	}
}
