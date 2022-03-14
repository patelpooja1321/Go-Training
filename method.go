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
	fmt.Println("Students Details:", std)

}

func (s person_details) Student() {
	fmt.Println("Students Status:", s.Status)
}

func (s *person_details) New_phone() {
	s.phone = 3456832
	fmt.Println("Students Status:", s.phone)
}
