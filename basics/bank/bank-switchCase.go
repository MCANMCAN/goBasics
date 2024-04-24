package main

import (
	"fmt"
	"os"
	"strconv"
	"errors"
)

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}
func getBalanceFromFile() (float64 , error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000 , errors.New("Failed to find balance file")
	}
	balanceText := string(data)
	balance , err := strconv.ParseFloat(balanceText,64)
	if err != nil {
		return 1000 , errors.New("Failed to parse balance value")
	}
	return balance, nil
}


func main() {

	var accountBalance, err = getBalanceFromFile()
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

		switch choise {

		case 1 :
			// accountBalance, _ = getBalanceFromFile()
			fmt.Println("Your Balance is", accountBalance)
		case 2 :
			fmt.Print("Your Deposit: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("The Deposit amount should be greater than 0!")
			
				// return
				continue
			}
			accountBalance += depositAmount
			writeBalanceToFile(accountBalance)
			fmt.Println("New Amount: ", accountBalance )

		case 3 : 
			fmt.Print("Enter the amount to Withdraw: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0  {
				fmt.Println("The Withdraw amount should be greater than 0!")
				// return
				continue
			} else if withdrawAmount >= accountBalance{
				fmt.Println("The Withdraw amount should not be greater than account balance!")
				// return
				continue
			}
			accountBalance -= withdrawAmount
			writeBalanceToFile(accountBalance)
			fmt.Println("New Amount: ", accountBalance )
		default : 
			fmt.Println("Exiting the Bank. Thank You! " )
			return
			// break // to break loop 

		}
	}
}