package main

import (
	"fmt"
	"time"
)

func main() {
	sample("direct")
	go sample("sub")
	go sample("third")

	time.Sleep(time.Second * 3)
	fmt.Println("end")
}

func sample(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s, ": ", i)
		time.Sleep(time.Second)
	}
}
