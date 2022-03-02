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

	createdate := time.Date(2021, time.February, 4, 03, 15, 2, 0, time.UTC)
	fmt.Println(createdate)
	fmt.Println(createdate.Format("01-04-2022 Monday"))

}
