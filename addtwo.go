package main

import "fmt"

var arr1 = []int{2, 4, 3}
var arr2 = []int{5, 6, 4}
var sum []int
var count int

func main() {
	fmt.Println("Add two array")
	l := len(arr1)
	for i := 0; i < l; i++ {
		fmt.Println(arr1[i], arr2[i])
		sum1 := arr1[i] + arr2[i]
		c := sum1 / 10
		if c == 1 {
			sum1 = sum1 + 1
			fmt.Println(sum)
		}
		sum = append(sum, sum1)
	}
	fmt.Println(sum)
}
