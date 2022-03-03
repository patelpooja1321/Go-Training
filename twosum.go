package main

import "fmt"

func main() {
	nums := []int{22, 7, 2, 15}
	target := 9
	l := len(nums)
	var sum int
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			sum = nums[i] + nums[j]
			if sum == target {
				fmt.Println(i, j)
			}
		}
	}

}
