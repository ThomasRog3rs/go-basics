package main

import (
	"fmt"
)

func main() {
	var currentOption int = 0
	balance := 0.00

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
			fmt.Println("Your balance: $", balance)
		case 2:
			depositAmount := 0.00
			fmt.Println("Current balance: $", balance)
			fmt.Print("How much to deposit?: ")
			fmt.Scan(&depositAmount)
			balance += depositAmount
		case 3:
			withdrawAmount := 0.00
			fmt.Println("Current balance: $", balance)
			fmt.Print("How much to withdraw?: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount > balance {
				fmt.Println("You don't have that much money, please select a smaller ammount!")
				break
			}
			balance = balance - withdrawAmount
		case 4:
			fmt.Println("Goodbye")
		default:
			fmt.Println("Invalid option")
		}

		fmt.Println("")
		fmt.Println("------------------------------------")
		fmt.Println("");
	}

}
