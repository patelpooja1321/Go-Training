package main

import "fmt"

func main() {
	price_list := []int{7, 1, 5, 3, 6, 4}
	max_dif, min_diff := 0, 0
	fmt.Println("Stock Price:", price_list)
	for i := 0; i <= len(price_list); i++ {
		for j := i + 1; j < len(price_list); j++ {
			price_dif := price_list[i] - price_list[j]
			//fmt.Println("Price diff ", price_dif)
			if price_dif > max_dif {
				max_dif = price_dif

			} else if price_dif < min_diff {
				min_diff = price_dif

			} else {
				continue
			}
		}
	}
	fmt.Println("Maximum Profile:", max_dif)
	fmt.Println("Min Stock Price", min_diff)
	// fmt.Println("Stock Purchase Price is", min_diff)
	// fmt.Println("Selling Price of stock is", max_dif)
	// fmt.Println("Current Profile is", max_dif-min_diff)

}
