// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	const inflationRate = 2.5
// 	var investmentAmount float64 = 1000
// 	expectedReturnRate := 5.5
// 	var years float64 = 10

// 	fmt.Print("Please input your investment amount: ")
// 	_, err := fmt.Scan(&investmentAmount)
// 	if err != nil {
// 		fmt.Println("Failed to read input, defaulting investement amount to 1000")
// 	}

// 	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
// 	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

// 	fmt.Println(futureValue)
// 	fmt.Println(futureRealValue)
// }

package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	revenue = getFloatInput("please enter revenue made: ")
	expenses = getFloatInput("Please enter your expenses: ")
	taxRate = getFloatInput("Please enter your taxRate: ")

	earningsBeforeTax, profit, ratio := doTheCalcs(revenue, expenses, taxRate)

	fmt.Printf("Earnings Before Tax (EBT): $%.2f\n", earningsBeforeTax)
	fmt.Printf("Profit After Tax: $%.2f\n", profit)
	fmt.Printf("EBT/Profit Ratio: %.2f\n", ratio)
}

func getFloatInput(message string) (userInput float64) {
	fmt.Print(message)
	_, err := fmt.Scan(&userInput)

	if err != nil {
		fmt.Println("Failed to read input")
		return 0.00
	}

	return
}

func doTheCalcs(revenue, expenses, taxRate float64) (earningsBeforeTax, profit, ratio float64) {
	earningsBeforeTax = revenue - expenses
	profit = earningsBeforeTax * (1 - (taxRate / 100))
	ratio = earningsBeforeTax / profit

	return
}
