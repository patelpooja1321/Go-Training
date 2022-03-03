package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Assignment 2 : Find Duplicate values from the Array")
	var arr1 []int
	var i, count int
	arr := []int{2, 3, 4, 5, 6, 6, 7, 7, 8, 8, 2, 4}
	l := len(arr)
	fmt.Println("length of Array:", l)
	fmt.Println("Orignal Array:", arr)
	sort.Ints(arr)
	fmt.Println("Sorted Array:", arr)
	for i = 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if arr[i] == arr[j] {
				arr1 = append(arr1, arr[i])
				count++
			}
		}
	}
	fmt.Println("Number of Duplicate data in Array:", count)
	fmt.Println("Duplicate data in Array is:", arr1)
}
