package main

import (
	"fmt"
)

func main() {
	var currentOption int = 0

	fmt.Println("Welcome to Go Bank!")

	for currentOption != 4 {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit App")

		fmt.Print("Please enter your option: ")
		fmt.Scan(&currentOption)

		switch currentOption {
		case 1:
			fmt.Println("Your balance")
		case 2:
			fmt.Println("Give us your money")
		case 3:
			fmt.Println("Here is your money")
		case 4:
			fmt.Println("Goodbye")
		default:
			fmt.Println("Invalid option")
		}
	}

}
