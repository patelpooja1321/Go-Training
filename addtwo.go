package main

import (
	"fmt"
)

var arr1 = []int{2, 4, 3, 5}
var arr2 = []int{5, 7, 4, 6}
var sum []int
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
			c = sum1 % 10
			sum = append(sum, c)
		}

	}
	l1 := len(sum)
	for i := l1; i > 0; i++ {
		//fmt.Println(sum)
	}

}
