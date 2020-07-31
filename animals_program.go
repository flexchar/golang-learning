package main

import (
	"fmt"
	"strings"
)

// Animal type
type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name, food, move, sound string
}

type Bird struct {
	name, food, move, sound string
}

type Snake struct {
	name, food, move, sound string
}

// Eat prints animal's food
func (a Cow) Eat() {
	fmt.Printf("%s (Cow) eats %s \n", a.name, a.food)
}

// Speak prints animal's sound
func (a Cow) Speak() {
	fmt.Printf("%s (Cow) says %s \n", a.name, a.sound)
}

// Move prints animal's move
func (a Cow) Move() {
	fmt.Printf("%s (Cow) moves %s \n", a.name, a.move)
}

// Eat prints animal's food
func (a Bird) Eat() {
	fmt.Printf("%s (Bird) eats %s \n", a.name, a.food)
}

// Speak prints animal's sound
func (a Bird) Speak() {
	fmt.Printf("%s (Bird) says %s \n", a.name, a.sound)
}

// Move prints animal's move
func (a Bird) Move() {
	fmt.Printf("%s (Bird) moves %s \n", a.name, a.move)
}

// Eat prints animal's food
func (a Snake) Eat() {
	fmt.Printf("%s (Snake) eats %s \n", a.name, a.food)
}

// Speak prints animal's sound
func (a Snake) Speak() {
	fmt.Printf("%s (Snake) says %s \n", a.name, a.sound)
}

// Move prints animal's move
func (a Snake) Move() {
	fmt.Printf("%s (Snake) moves %s \n", a.name, a.move)
}

func main() {

	animals := make(map[string]Animal)

	for {
		fmt.Print("> ")

		var name string
		var query string
		var action string // action can also be a type of animal
		_, _ = fmt.Scanf("%s %s %s", &query, &name, &action)

		name = strings.Trim(name, " ")
		query = strings.Trim(query, " ")
		action = strings.Trim(action, " ")

		// Detect program action
		switch query {
		case "newanimal":

			if name == "" || action == "" {
				fmt.Println("Please type QUERY NAME ACTION||TYPE")
				continue
			}

			// Detect animal type (name)
			switch action {
			case "cow":
				animals[name] = Cow{
					name:  name,
					food:  "grass",
					move:  "walk",
					sound: "moo",
				}

			case "bird":

				animals[name] = Bird{
					name:  name,
					food:  "worms",
					move:  "fly",
					sound: "peep",
				}

			case "snake":
				animals[name] = Snake{
					name:  name,
					food:  "grass",
					move:  "slither",
					sound: "hss",
				}
			}

			fmt.Println("Created it!")

		case "query":

			animal, isFound := animals[name]

			if isFound == false {
				fmt.Println("Couldn't find such animal, sorry bud")
				continue
			}

			switch action {
			case "eat":
				animal.Eat()

			case "move":
				animal.Move()

			case "speak":
				animal.Speak()

			default:
				fmt.Println("Animal has not evolved to do this much just yet ;)")
			}

		default:
			fmt.Println("Such command does not exist")
		}

	}

}
