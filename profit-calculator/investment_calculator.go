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

	fmt.Print("Please enter revenue made: ")
	fmt.Scan(&revenue)

	fmt.Print("Please enter your expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Please enter your taxRate: ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - (taxRate / 100))

	ratio := earningsBeforeTax / profit

	fmt.Printf("Earnings Before Tax (EBT): $%.2f\n", earningsBeforeTax)
	fmt.Printf("Profit After Tax: $%.2f\n", profit)
	fmt.Printf("EBT/Profit Ratio: %.2f\n", ratio)
}
