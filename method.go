package main

import "fmt"

type person_details struct {
	Name   string
	id_no  int
	Status bool
	phone  int
}

func main() {
	std := person_details{"Meera", 12345, true, 12345678}
	fmt.Println("Students Details:", std)
	std.Student()

}

func (s person_details) Student() {
	fmt.Println("Students Status:", s.Status)
}
