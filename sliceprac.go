package main

import "fmt"

func main() {
	fmt.Println("Slice Practice")
	var arr = []string{"abc", "pqr", "xyz", "mno", "lmr"}
	fmt.Println("Orginal Array is", arr)
	var i int = 3
	arr = append(arr[:i])
	fmt.Println("Slice is", arr)
	newname := make([]string, 3)
	newname[0] = "patel"
	newname[1] = "Pooja"
	newname[2] = "Dhaval"
	fmt.Println("New Slice:", newname)

	newname = append(newname, "Mihir", "Dhruvil")
	fmt.Println("Slice updated:", newname)
	newname = append(newname[:i], newname[i+1:]...)
	fmt.Println("Remove element from slice:", newname)
}
