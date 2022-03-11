package main

import "fmt"

type circle struct {
	Pi float32
	r  int
}

func main() {
	fmt.Println("Find Area of Circle:")
	circle := circle{13.4, 8}
	Area := circle.Pi * float32(circle.r)
	fmt.Println("Circle Area is:", Area)

}
