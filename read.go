// Write a program which reads information from a file and represents it in a slice of structs. Assume that there is a text file which contains a series of names. Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

// Your program will define a name struct which has two fields, fname for the first name, and lname for the last name. Each field will be a string of size 20 (characters).

// Your program should prompt the user for the name of the text file. Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file. Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file. After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type person struct {
	fname string
	lname string
}

func main() {
	listOfNames := make([]person, 0, 0)

	var fileName string
	fmt.Println("Enter file name")
	_, _ = fmt.Scanln(&fileName)

	contents, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
	}

	names := strings.Split(string(contents), "\n")

	for _, name := range names {
		parts := strings.Fields(name)

		theName := person{
			fname: parts[0],
			lname: parts[1],
		}
		listOfNames = append(listOfNames, theName)

		// fmt.Println(parts)
	}

	for _, name := range listOfNames {
		fmt.Printf("First name: %s, Last name: %s\n", name.fname, name.lname)
	}

}
