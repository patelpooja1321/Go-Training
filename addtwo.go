package main

import "fmt"

var arr1 = []int{2, 4, 3}
var arr2 = []int{5, 6, 4}
var sum []int
var count int
var c int

func main() {
	fmt.Println("Add two array")
	fmt.Println("First Array:", arr1)
	fmt.Println("Second Array:", arr2)
	l := len(arr1) - 1
	for i := l; i >= 0; i-- {
		sum1 := arr1[i] + arr2[i]

		if c == 1 {
			sum1 = sum1 + 1
			sum = append(sum, sum1)
			c = 0
		} else {
			c1 := sum1 % 10
			sum = append(sum, c1)
			c = sum1 / 10
		}

	}
	fmt.Println(sum)
}
