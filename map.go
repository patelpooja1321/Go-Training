package main

import "fmt"

func main() {
	lauguage := make(map[int]string)

	lauguage[1] = "JS"
	lauguage[2] = "Python"
	lauguage[3] = "Java"
	lauguage[4] = "C++"

	fmt.Println(lauguage)
	delete(lauguage, 3)
	fmt.Println(lauguage[4])
	fmt.Println(lauguage)

	for k, v := range lauguage {
		fmt.Println("Key is %v and value is %v", k, v)
	}

}
