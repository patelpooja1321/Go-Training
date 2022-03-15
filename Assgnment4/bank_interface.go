package main

import "fmt"

type Account struct {
	Total_balance int
	Amount        int
}

func (a *Account) Add_money() int {
	fmt.Println("Please Enter Amount For deposit:")
	fmt.Scan(&a.Amount)
	a.Total_balance = a.Total_balance + a.Amount
	//fmt.Println("Current balance is:", a.Total_balance)
	return a.Total_balance
}

func (a *Account) withdraw_money() int {
	fmt.Println("Please Enter Amount for Withdraw:")
	fmt.Scan(&a.Amount)
	return a.Total_balance - a.Amount

}

type Account_interface interface {
	Add_money()
	withdraw_money()
}

func main() {
	fmt.Println("Welcome to online Bank system")
	var choice int
	var i int = 0
	var user = Account{Total_balance: 50}
	//fmt.Print("Account Holder Details:", user)
	fmt.Print("For add Money in your Account please press: 1, for withdraw money enter: 2 and for Exit: 0= ")
	fmt.Scan(&choice)
	for i == 0 {
		switch choice {
		case 1:
			fmt.Println("Current balance is:", user.Add_money())
			fmt.Print("For add Money in your Account please press: 1, for withdraw money enter: 2 and for Exit: 0= ")
			fmt.Scan(&choice)
		case 2:
			fmt.Println("Current balance is:", user.withdraw_money())
			fmt.Print("For add Money in your Account please press: 1, for withdraw money enter: 2 and for Exit: 0= ")
			fmt.Scan(&choice)
		default:
			i = 1
			fmt.Println("Thank You For using our bank system!")
		}
	}

}
