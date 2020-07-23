// Write a program which prompts the user to first enter a name, and then enter an address. Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively. Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Hey there!")

	var input string

	user := make(map[string]string)

	fmt.Println("Type your name")
	fmt.Scanln(&input)

	user["name"] = input

	fmt.Println("Type your address")
	fmt.Scanln(&input)

	user["address"] = input

	json, _ := json.MarshalIndent(user, "", "\t")

	fmt.Printf("\n%s\n", json)

}
