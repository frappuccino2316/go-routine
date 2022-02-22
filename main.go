package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go receiver(ch)

	i := 0
	for i < 100 {
		ch <- i
		i++
	}

	time.Sleep(time.Second)
}

func receiver(ch chan int) {
	for {
		i := <-ch
		fmt.Println(i)
	}
}
