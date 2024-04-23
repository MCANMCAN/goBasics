package main

import "fmt"

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
	return userInput

}
