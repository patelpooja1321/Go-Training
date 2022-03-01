package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Please, Enter your name:")
	var name string
	fmt.Scanln(&name)
	rand.Seed(20)
	fmt.Println("random number is", rand.Intn(10))
}
