package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64 = 1000
	expectedReturnRate := 5.5
	var years float64 = 10

	fmt.Print("Please input your investment amount: ")
	_, err := fmt.Scan(&investmentAmount)
	if err != nil {
		fmt.Println("Failed to read input, defaulting investement amount to 1000")
	}

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
