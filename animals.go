package main

import (
	"fmt"
	"strings"
)

// Animal type
type Animal struct {
	food, move, sound string
}

// Eat prints animal's food
func (a Animal) Eat(name string) {
	fmt.Println(name + " eats " + a.food)
}

// Speak prints animal's sound
func (a Animal) Speak(name string) {
	fmt.Println(name + " says " + a.sound)
}

// Move prints animal's move
func (a Animal) Move(name string) {
	fmt.Println(name + " moves " + a.move)
}

func main() {
	cow := Animal{
		food:  "grass",
		move:  "walk",
		sound: "moo",
	}

	bird := Animal{
		food:  "worms",
		move:  "fly",
		sound: "peep",
	}

	snake := Animal{
		food:  "grass",
		move:  "slither",
		sound: "hss",
	}

	animals := make(map[string]Animal)
	animals["cow"] = cow
	animals["bird"] = bird
	animals["snake"] = snake

	for {
		fmt.Print("> ")

		var name string
		var action string
		_, _ = fmt.Scan(&name, &action)

		name = strings.Trim(name, " ")
		action = strings.Trim(action, " ")

		if name == "" || action == "" {
			fmt.Println("Something's wrong, please type ANIMAL (space) ACTION")
			return
		}

		animal, isFound := animals[name]

		if isFound == false {
			fmt.Println("Couldn't find such animal, sorry bud")
			return
		}

		switch action {
		case "eat":
			animal.Eat(name)

		case "move":
			animal.Move(name)

		case "speak":
			animal.Speak(name)

		default:
			fmt.Println("Animal has not evolved to do this much just yet ;)")
		}
	}

}
