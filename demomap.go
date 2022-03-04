package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in Golang")
	lang := make(map[int]string)
	lang[1] = "Java"
	lang[2] = "C++"
	lang[3] = "Golang"

	fmt.Println("map list: ", lang)
	fmt.Println("Get key 2 value from a list: ", lang[2])

	delete(lang, 2)
	fmt.Println("map list: ", lang)

	for key, value := range lang {
		fmt.Printf("In map key is %v and value is %v \n", key, value)
	}
}
