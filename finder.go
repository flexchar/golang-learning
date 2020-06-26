package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Println("Yoo! enter something ")

	fmt.Scan(&input)

	fmt.Printf("You typed: %s \n", input)

	// The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.
	normalized := strings.ToLower(input)

	if strings.HasPrefix(normalized, "i") &&
		strings.HasSuffix(normalized, "n") &&
		strings.Contains(normalized, "a") {
		// The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’.

		fmt.Println("Found!")
	} else {
		// The program should print “Not Found!” otherwise.
		fmt.Println("Not Found!")

	}

}
