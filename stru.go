package main

import "fmt"

func main() {

	fmt.Println("Struct in Golang")
	Pooja := employee{"Pooja", 1233459877, "Montgomery", "GolangDeveloper", 20, true}
	fmt.Println("Pooja Details:", Pooja)

}

type employee struct {
	Name         string
	Phone_Number int
	City         string
	Position     string
	Age          int
	OPT_status   bool
}
