package main

import (
	"fmt"
)

func main() {
	var floatNum float32
	fmt.Println("Hi there!")

	// which prompts the user to enter a floating point number
	fmt.Println("Type float...")

	_, _ = fmt.Scan(&floatNum)

	// and prints the integer which is a truncated version of the floating point number that was entered.
	fmt.Printf("the input is: %d \n", uint(floatNum))
}
