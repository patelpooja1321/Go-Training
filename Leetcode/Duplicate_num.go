package main

import (
	"fmt"
)

func main() {
	fmt.Println("Duplicate Number!")
	num := []int{1, 2, 3, 1}
	var result bool
	//sort.Ints(num)
	fmt.Println("Orignal Data:", num)
	for i := 0; i < len(num); i++ {
		for j := i + 1; j < len(num); j++ {
			//fmt.Println(num[i], num[j])
			if num[i] == num[j] {
				result = true
				break
			} else {
				result = false
			}

		}

	}
	fmt.Println(result)

}
