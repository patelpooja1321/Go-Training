package main

import (
	"fmt"
	"strings"
)

var sentence string = "this is java and java is good"
var i, j int
var arr []string

func main() {
	fmt.Println("Count the number of occurrences of a word in a sentence")
	fmt.Println("Statment is:", sentence)
	count := 0
	arr = strings.Split(sentence, " ")
	l := len(arr)
	for i = 0; i < l; i++ {
		for j = i + 1; j < l; j++ {
			if arr[i] == arr[j] {
				count++
				fmt.Printf("Same Word in sentence: %v \n index number is %v %v \n", arr[i], i, j)
			}
		}
	}

	fmt.Println("Number of occurrence of word in sentence is", count)
}
