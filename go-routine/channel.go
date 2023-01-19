package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	// 2. Buffered Channel
	messages := make(chan int, 3)

	go func() {
		for {
			i := <-messages
			fmt.Println("Receive Data", i)
		}
	}()
	for i := 0; i < 5; i++ {
		fmt.Println("Send data", i)
		messages <- i
	}
	// 1. channel
	// messages := make(chan string)

	// sayHelloTo := func(who string) {
	// 	data := fmt.Sprintf("Hello %s", who)
	// 	messages <- data
	// }
	// go sayHelloTo("Arif1")
	// go sayHelloTo("Arif3")
	// go sayHelloTo("Arif2")

	// message1 := <-messages
	// fmt.Println(message1)

	// message2 := <-messages
	// fmt.Println(message2)

	// message3 := <-messages
	// fmt.Println(message3)

}
