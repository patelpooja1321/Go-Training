package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Welcome to time Class")
	currenttime := time.Now()
	fmt.Println(currenttime)
	fmt.Println("Time Formate:", currenttime.Format("02-01-2022 Monday 15:04:05 PM"))

}
