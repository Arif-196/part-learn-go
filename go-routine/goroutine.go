package go_routine

import (
	"fmt"
	"runtime"
)

func RunHelloWorld() {
	fmt.Println("Hello World!!!")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		DisplayNumber(i)
	}
	time.Sleep(30 * time.Second)
}

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

