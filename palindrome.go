package main

import "fmt"

func main() {
	var num string = "1211"
	fmt.Println("palindrome Number is", num)
	l := len(num)
	l1 := l / 2
	//j := l/2

	for i := 0; i < l1; i++ {
		for j := l1; j < l; j++ {
			if num[i] == num[j] {
				fmt.Println("Number is Palindrome")
			} else {
				fmt.Println("Number is not Palindrome")
				break
			}
		}

	}

}
