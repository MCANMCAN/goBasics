package main

import (
	"fmt"

	"example.com/bank/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {

	var accountBalance, err = fileops.GetFloatFromFile("accountBalanceFile")
	if err != nil {
		fmt.Println("Can not get Balance value.")
		fmt.Println(err)
		panic(err)
	}
	var depositAmount float64
	var withdrawAmount float64

	// for i :=0 ; i < 2; i++ { -> finite loop
	fmt.Println("welcome th Go Bank")
	for {

		presentOptions()

		var choise int
		fmt.Print("Choise an operation:")
		fmt.Scan(&choise)
		fmt.Println("Your Choise : ", choise)

		switch choise {

		case 1:
			// accountBalance, _ = getBalanceFromFile()
			fmt.Println("Your Balance is", accountBalance)
		case 2:
			fmt.Print("Your Deposit: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("The Deposit amount should be greater than 0!")

				// return
				continue
			}
			accountBalance += depositAmount
			fileops.WriteValueToFile(accountBalance, accountBalanceFile)
			fmt.Println("New Amount: ", accountBalance)

		case 3:
			fmt.Print("Enter the amount to Withdraw: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("The Withdraw amount should be greater than 0!")
				// return
				continue
			} else if withdrawAmount >= accountBalance {
				fmt.Println("The Withdraw amount should not be greater than account balance!")
				// return
				continue
			}
			accountBalance -= withdrawAmount
			fileops.WriteValueToFile(accountBalance, accountBalanceFile)
			fmt.Println("New Amount: ", accountBalance)
		default:
			fmt.Println("Exiting the Bank. Thank You! ")
			return
			// break // to break loop

		}
	}
}
