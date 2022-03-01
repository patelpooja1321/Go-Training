package main

import "fmt"

var msg = "od Morning!"

const Loginsession = true //Public data type

func main() {
	fmt.Println("Message is:", msg)
	var name string = "Pooja"
	fmt.Println("My name is", name)
	fmt.Printf("Variable name Value is %s \n", name)
	fmt.Printf("Variable name data type is %T \n", name)

	var choice bool = true
	fmt.Printf("Variable choice data type is %T\n", choice)

	var smallval uint8 = 255 //range is 0 to 255
	fmt.Println("variable value is", smallval)
	fmt.Printf("Variable data type is %T \n", smallval)

	var smallint int = 1255
	fmt.Println("variable value is", smallint)
	fmt.Printf("Variable data type is %T \n", smallint)

	var smallfloat float32 = 255.123456789
	fmt.Println("variable value is", smallfloat)
	fmt.Printf("Variable data type is %T \n", smallfloat)

	var smallfloat64 float64 = 255.12345678901234567890
	fmt.Println("variable value is", smallfloat64)
	fmt.Printf("Variable data type is %T \n", smallfloat64)

	//Default value of variable
	var defaultval bool
	fmt.Println(defaultval)
	fmt.Printf("Default value of variable data type is %T \n", defaultval)

	//implicit Type
	var color = "Blue"
	fmt.Printf("Variable color data type is %T \n", color)

	//No var type
	luckynumber := 4.0
	fmt.Printf("Variable luckynumber data type is %T \n", luckynumber)

	fmt.Println(Loginsession)
	fmt.Printf("Variable Loginsession data type is %T \n", Loginsession)

}
