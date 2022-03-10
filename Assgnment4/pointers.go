package main

import "fmt"

func main() {
	a := 255
	val := 0
	var b *int
	if b == nil {
		fmt.Println("Nill Pointer:", b)
	}
	b = &a
	var ptr **int = &b
	fmt.Printf("Type of b is %T\n", b)
	fmt.Println("address of a is", b)
	fmt.Println("address of b is", ptr) //Double pointer
	if val == 0 {
		panic("Pointer Arithmetic is not support in Golang")
	}

}
