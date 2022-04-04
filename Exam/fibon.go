//package main

// import "fmt"

// func fibonacci(c, quit chan int) {
// 	x, y := 0, 1
// 	for {
// 		select {
// 		case c <- x:
// 			x, y = y, x+y
// 		case <-quit:
// 			fmt.Println("quit")
// 			return
// 		}
// 	}
// }

// func main() {
// 	c := make(chan int)
// 	quit := make(chan int)
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			fmt.Println(<-c)

// 		}
// 		quit <- 0
// 	}()
// 	fibonacci(c, quit)
// }

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var i, j int
var arr []string

func main() {
	sentence, err := ioutil.ReadFile("./data.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sentence)
	count := 0
	arr = strings.Split(sentence, " ")
	l := len(arr)
	for i = 0; i < l; i++ {
		for j = i + 1; j < l; j++ {
			if arr[i] == arr[j] {
				count++
				fmt.Printf("Same Word in sentence: %v \n index number is %v %v \n", arr[i], i, j)
			}
		}
	}

	fmt.Println("Number of occurrence of word in sentence is", count)

}
