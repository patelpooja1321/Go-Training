package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	msg := "Practice of user input"
	fmt.Println(msg)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating for Pizza:")

	//comma or || error or
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)

}
