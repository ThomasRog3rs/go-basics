package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func writeBalanceToFile(balance float64) {
	balanceTxt := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceTxt), 0644)
}

func getCurrentBalance() float64 {
	balance := 0.00

	file, err := os.Open("balance.txt")
	if err != nil {
		log.Fatal(err)
		fmt.Printf("Failed to read balance!")
		return balance
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		parsedValue, parseError := strconv.ParseFloat(scanner.Text(), 64)
		if parseError != nil {
			log.Fatal("Cannot read line")
			balance = 0.00
		} else {
			balance = parsedValue
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		fmt.Printf("Failed to read balance!")
		balance = 0.00
	}

	return balance
}

func updateBalance(amount float64) {
	currentBalance := getCurrentBalance()
	currentBalance += amount
	writeBalanceToFile(currentBalance)
}

func main() {
	var currentOption int = 0
	// balance := 0.00

	fmt.Println("Welcome to Go Bank!")

	for currentOption != 4 {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit App")

		fmt.Print("Please enter your option: ")
		fmt.Scan(&currentOption)

		fmt.Println("")
		fmt.Println("-----------Your Choice:-------------")
		fmt.Println("")

		switch currentOption {
		case 1:
			fmt.Println("Your balance: $", getCurrentBalance())
		case 2:
			depositAmount := 0.00
			fmt.Println("Current balance: $", getCurrentBalance())
			fmt.Print("How much to deposit?: ")
			fmt.Scan(&depositAmount)
			if depositAmount < 0 {
				fmt.Println("You can only deposit with posative values, obviously...")
				break
			}

			updateBalance(depositAmount)
		case 3:
			currentBalance := getCurrentBalance()
			withdrawAmount := 0.00
			fmt.Println("Current balance: $", currentBalance)
			fmt.Print("How much to withdraw?: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount > currentBalance {
				fmt.Println("You don't have that much money, please select a smaller ammount!")
				break
			}

			if withdrawAmount > 0 {
				withdrawAmount = -withdrawAmount
			}
			updateBalance(withdrawAmount)
		case 4:
			fmt.Println("Goodbye")
		default:
			fmt.Println("Invalid option")
		}

		fmt.Println("")
		fmt.Println("------------------------------------")
		fmt.Println("")
	}

}
