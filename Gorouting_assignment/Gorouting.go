package main

import (
	"fmt"
)

func sqare(n int, c, quit chan int) {
	//x, y := 0, 1
	//num := 0
	for i := 0; i < n; i++ {

		select {
		case c <- i:
			//x, y = x*x, x+1
			c <- i * i
		case <-quit:
			fmt.Println("quit")
			return
		}
	}

}

func main() {
	c := make(chan int, 10)
	quit := make(chan int)
	go sqare(cap(c), c, quit)
	for i := range c {
		fmt.Println(i)
	}
	close(c)
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	quit := make(chan int)
// 	c := generator()
// 	receiver(c, quit)
// }

// func receiver(c, quit <-chan int) {
// 	for v := range c {
// 		select {
// 		case s := <-c:
// 			s = v * v
// 			fmt.Println(s)
// 		case <-quit:
// 			return
// 		}

// 	}
// }

// func generator() <-chan int {
// 	c := make(chan int)
// 	//quit := make(chan int)
// 	go func() {
// 		for i := 0; i < 11; i++ {
// 			c <- i
// 		}
// 		close(c)
// 	}()

// 	return c
// }
