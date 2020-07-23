package main

import (
	"fmt"
	"sort"
	"strconv"
)

// Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
func main() {
	fmt.Println("Hi there!")

	// create empty integer slice of size (length) 3
	theSlice := make([]int, 3)
	exitChar := "X"
	var input string

	for {
		// During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
		fmt.Println("Type int...")

		_, _ = fmt.Scanln(&input)
		// fmt.Printf("\ninput: %s", input)

		// The slice must grow in size to accommodate any number of integers which the user decides to enter. The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.
		if input == exitChar {
			break
		}

		theInt, _ := strconv.Atoi(input)

		// The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
		theSlice = append(theSlice, theInt)
		sort.Slice(theSlice, func(i, j int) bool {
			return theSlice[i] < theSlice[j]
		})

		fmt.Println(theSlice)

	}
	fmt.Println("Goodbye")
}
