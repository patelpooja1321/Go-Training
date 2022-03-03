package main

import "fmt"

func main() {

	fmt.Println("Wellcome in  Array Class")
	var arr [4]int
	arr[0] = 1
	arr[1] = 2
	arr[3] = 4
	fmt.Println("Value of array:", arr)
	fmt.Println("Value of array:", len(arr))

	var arr1 = [5]string{"abc", "pqr", "xyz", ""}
	fmt.Println("Value of array:", arr1)
	fmt.Println("Value of array:", len(arr1))

}
