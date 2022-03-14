package main

import "fmt"

type circle struct {
	Pi float32
	r  int
}

func main() {
	fmt.Println("Find Area of Circle:")
	c := circle{}
	c.Pi = 13.4
	c.r = 8
	Area := c.Pi * float32(c.r)
	fmt.Println("Circle Area is:", Area)

}
