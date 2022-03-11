package main

import "fmt"

type circle struct {
	Pi float32
	r  int
}

func main() {
	fmt.Println("Find Area of Circle:")
	c := circle{13.4, 8}
	Area := c.Pi * float32(c.r)
	fmt.Println("Circle Area is:", Area)

}
