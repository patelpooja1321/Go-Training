package main

import "fmt"

func main() {
	fmt.Println("Wellcome in pointer class")
	var ptr *int
	fmt.Println("initilize value of pointer", ptr)

	myvalue := 13
	var ptr1 = &myvalue
	fmt.Println("Value of Pointer", ptr1)
	fmt.Println("location point by Pointer", *ptr1)

	*ptr1 = *ptr1 + 2
	fmt.Println(" New Value of Pointer", myvalue)

}
