package main

import (
	"fmt"
)

func main() {
	var accountBalance = 1000.0
	var depositAmount float64
	var withdrawAmount float64
	fmt.Println("welcome th Go Bank")
	fmt.Println("Choose an Operation!")
	fmt.Println(
		`
		1. Check Balance 
		2. Deposit Money
		3. Withdraw Money
		4. Exit
		`)
	
	var choise int 
	fmt.Print("Choise an operation:" )
	fmt.Scan(&choise)
	fmt.Println("Your Choise : ", choise)


	if choise == 1 {
		fmt.Println("Your Balance is", accountBalance)
	} else if choise == 2 {
		fmt.Print("Your Deposit: ")
		fmt.Scan(&depositAmount)
		accountBalance += depositAmount
		fmt.Println("New Amount: ", accountBalance )
	} else if choise == 3 {
		fmt.Print("Enter the amount to Withdraw: ")
		fmt.Scan(&withdrawAmount)
		accountBalance -= withdrawAmount
		fmt.Println("New Amount: ", accountBalance )
	} else {
		fmt.Println("Exiting the Bank. Thank You! " )
	}
}