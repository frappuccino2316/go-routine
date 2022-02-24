package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	ch2 := make(chan int)

	go receiver(ch)
	go receiver(ch2)

	fmt.Println("first")
	for i := range ch {
		fmt.Println(i)
	}

	fmt.Println("second")
	for i := range ch2 {
		fmt.Println(i)
	}

	time.Sleep(time.Second)
}

func receiver(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
}
