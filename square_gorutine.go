package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int)
	defer close(c)

	// sender
	go func() {
		for i := 0; i <= 10; i++ {
			c <- i * i
		}
	}()

	// receiver
	go func() {
		for v := range c {
			fmt.Println(v)
		}
	}()

	time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	fmt.Println("Done")
}
