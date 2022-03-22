// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func main() {
// 	c := make(chan int)
// 	defer close(c)

// 	// sender
// 	go func() {
// 		for i := 0; i <= 10; i++ {
// 			c <- i * i
// 		}
// 	}()

// 	// receiver
// 	go func() {
// 		for v := range c {
// 			fmt.Println(v)
// 		}
// 	}()

// 	time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
// 	fmt.Println("Done")
// }

package main

import (
	"fmt"
)

func main() {
	Numbers := 100
	// generate the common channel with inputs
	inputChan := generatePipeline(Numbers)

	// Fan-out to 2 Go-routine
	c1 := squareNumber(inputChan)
	//c2 := squareNumber(inputChan)

	// Fan-in the resulting squared numbers
	//c := fanIn(c1, c2)
	//sum := 0

	// Do the summation
	// for i := 0; i < Numbers; i++ {
	// 	sum = <-c1 * <-c2
	// }
	fmt.Printf("Total Sum of Squares: %d", c1)
}

func generatePipeline(numbers int) <-chan int {
	out := make(chan int)
	go func() {
		// for _, n := range numbers {
		for n := 1; n < numbers; n++ {

			out <- n
		}
		close(out)
	}()
	return out
}

func squareNumber(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func fanIn(input1, input2 <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}
