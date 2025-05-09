package main

import "fmt"

func main() {
	age := 32 // regular variable

	agePointer := &age // pointer variable

	fmt.Println("Age: ", *agePointer)

	adultYears := getAdultYears(agePointer)
	fmt.Println(adultYears)
}

func getAdultYears(age *int) int {
	return *age - 18
}