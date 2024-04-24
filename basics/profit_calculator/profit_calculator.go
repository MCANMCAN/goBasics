package main

import "fmt"
import "errors"
import "os"
// import "strconv"

const resultFile = "results.txt"

func main() {

	var revenue float64
	var expenses float64
	var tax_rate float64
	var earningsBT float64
	var earningsAT float64
	var ratio float64

	fmt.Println("@@@@ Profit Calculator @@@")
	// Receive user input for 'Revenue' as float64
	// fmt.Print("Please Enter the value of Revenue: ")
	// fmt.Scan(&revenue)
	revenue = getUserInfo("Please Enter the value of Revenue: ")

	// Receive user input for 'Expenses' as float64
	// fmt.Print("Please Enter the value of Expenses: ")
	// fmt.Scan(&expenses)
	expenses = getUserInfo("Please Enter the value of expenses: ")
	// Receive user input for 'Tax Rate' as float64
	// fmt.Print("Please Enter the value of VAT: ")
	// fmt.Scan(&tax_rate)
	tax_rate = getUserInfo("Please Enter the value of tax_rate: ")

	earningsBT, earningsAT, ratio = calculate(revenue, expenses, tax_rate)

	full_text := fmt.Sprintf(
`Earnings Before Tax = %.3f
Earnings After Tax = %.3f
The Ratio = %.3f `,earningsBT,earningsAT,ratio)

	os.WriteFile(resultFile, []byte(full_text),444)
	// fmt.Println(full_text)
	// func writeBalanceToFile(balance float64) {
	// 	balanceText := fmt.Sprint(balance)
	// 	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)	

	fmt.Println("@@@@ RESULTS @@@")
	fmt.Print("Earnings before Tax= ")
	fmt.Println(earningsBT)
	fmt.Println("------------")
	fmt.Print("Earnings after Tax= ")
	fmt.Println(earningsAT)
	fmt.Println("------------")
	fmt.Printf("Ratio= %.3f \n", ratio)
	fmt.Println("@@@@ Profit Calculator @@@")
}

func calculate(revenue, expenses, tax_rate float64) (float64, float64, float64) {
	EBT := revenue - expenses
	EAT := EBT * (1 - tax_rate/100)
	rt := EBT / EAT
	return EBT, EAT, rt
}

func getUserInfo(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	// fmt.Print("Please Enter the value of Revenue: ")
	fmt.Scan(&userInput)
	_ , err := checkInputValidity(userInput)
	if err != nil {
		fmt.Println(err)
		panic("the input is not valid:\n")
	}
	return userInput

}

func checkInputValidity(Uinput float64) (bool,error) {

		if Uinput <=0{
			return false ,  errors.New("Input cannot be equal or lower than 0! ")
		}
	return true , nil
}