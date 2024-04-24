package main

import (
	"fmt"
)

func main() {
	var accountBalance = 1000.0
	var depositAmount float64
	var withdrawAmount float64

	// for i :=0 ; i < 2; i++ { -> finite loop
		fmt.Println("welcome th Go Bank")
	for {

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
			if depositAmount <= 0 {
				fmt.Println("The Deposit amount should be greater than 0!")
				// return
				continue
			}
			accountBalance += depositAmount
			fmt.Println("New Amount: ", accountBalance )
		} else if choise == 3 {
			fmt.Print("Enter the amount to Withdraw: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0  {
				fmt.Println("The Withdraw amount should be greater than 0!")
				// return
				continue
			} else if withdrawAmount <= accountBalance{
				fmt.Println("The Withdraw amount should not be greater than account balance!")
				// return
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Println("New Amount: ", accountBalance )
		} else {
			fmt.Println("Exiting the Bank. Thank You! " )
			// return
			break // to break loop
		}
	}
	fmt.Println("Exiting the Bank. Thank You! " )
}