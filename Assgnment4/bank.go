package main

import "fmt"

func main() {
	fmt.Println("Welcome in abc Bank")
	var name string
	var account_number int
	var account_type string
	var total_balance int = 50
	var add_money int
	var withdraw_money int
	var choice int

	fmt.Print("Please, Enter Account Holder name:")
	fmt.Scanln(&name)
	fmt.Print("Enter Account Number:")
	fmt.Scanln(&account_number)
	fmt.Print("Please Enter Account type: Saving or Current : ")
	fmt.Scanln(&account_type)
	fmt.Print("For add Money in your Account please press 1 and for withdraw money enter 2:")
	fmt.Scan(&choice)
	defer fmt.Println("last balnce is:", total_balance)

	for {
		switch choice {
		case 1:
			fmt.Println("Please Enter Amount For deposit:")
			fmt.Scan(&add_money)
			total_balance = total_balance + add_money
			fmt.Println("Current balance is:", total_balance)
		case 2:
			fmt.Println("Please Enter Amount for Withdraw:")
			fmt.Scan(&withdraw_money)
			total_balance = total_balance - withdraw_money
			fmt.Println("Current balance is:", total_balance)
			if total_balance == 0 {
				panic("Balance is 0, Plase diposit money in next 15days")
			}
		case ' ':
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recover the Error", r)
					fmt.Print("For add Money in your Account please press 1 and for withdraw money enter 2:")
					fmt.Scan(&choice)
				}
			}()
			panic("Please Enter correct choice for bank operations")
		default:
			fmt.Println("Thank You For using our bank system!")
		}
		fmt.Print("For add Money in your Account please press 1 and for withdraw money enter 2:")
		fmt.Scan(&choice)
	}

}
