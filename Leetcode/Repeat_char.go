package main

import "fmt"

func main() {
	s := []string{"a", "b", "a", "b", "c"}
	count := 0
	var rept []string
	fmt.Println("String is", s)
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				rept = append(rept, s[i])
				count++
			}
		}

	}
	fmt.Printf("The Answer is %s , with the length of %v", rept, count)
}
