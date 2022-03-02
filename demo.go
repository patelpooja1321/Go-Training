package main

import (
	"fmt"
	"sort"
)

func input(x []int, err error) []int {
	if err != nil {
		return x
	}
	var d int
	n, err := fmt.Scanf("%d", &d)
	if n == 1 {
		x = append(x, d)
	}
	return input(x, err)
}
func main() {
	fmt.Println("Welcome to sope factory")
	var q int
	fmt.Println("Please, Enter the amount of soup that you want to supply:")
	fmt.Scanln(&q)
	fmt.Println("Enter the soap size")
	x := input([]int{}, nil)
	fmt.Println("Input:", x)
	sort.Ints(x)
}
