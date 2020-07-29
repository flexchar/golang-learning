package main

import (
	"fmt"
	"math"
)

// For example, let’s say that I want to assume the following values for acceleration, initial velocity, and initial displacement: a = 10, vo = 2, so = 1. I can use the following statement to call GenDisplaceFn() to generate a function fn which will compute displacement as a function of time.

// You will need to define and use a function called GenDisplaceFn() which takes three float64 arguments, acceleration a, initial velocity vo, and initial displacement so. GenDisplaceFn() should return a function which computes displacement as a function of time, assuming the given values acceleration, initial velocity, and initial displacement. The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one float64 argument which is the displacement travelled after time t.

// GenDisplaceFn will generate a function fn which will compute displacement as a function of time.
func GenDisplaceFn(acceleration, velocity, sec float64) func(float64) float64 {
	fn := func(time float64) float64 {
		// s =½ a t2 + vot + so
		return (0.5 * acceleration * math.Pow(time, 2)) + (velocity * time) + sec
	}
	return fn
}

func main() {
	// Let us assume the following formula for displacement s as a function of time t, acceleration a, initial velocity vo, and initial displacement so.

	// Write a program which first prompts the user to enter values for acceleration, initial velocity, and initial displacement. Then the program should prompt the user to enter a value for time and the program should compute the displacement after the entered time.

	var acceleration,
		velocity,
		displacement,
		timeToWait float64

	fmt.Println("Type acceleration...")
	fmt.Scan(&acceleration)

	fmt.Println("Type velocity...")
	fmt.Scan(&velocity)

	fmt.Println("Type displacement...")
	fmt.Scan(&displacement)

	// fn := GenDisplaceFn(10, 2, 1)
	fn := GenDisplaceFn(acceleration, velocity, displacement)

	// Then I can use the following statement to print the displacement after 3 seconds.
	// fmt.Println(fn(3))

	fmt.Println("Type the time to wait...")
	fmt.Scan(&timeToWait)

	// And I can use the following statement to print the displacement after 5 seconds.
	// fmt.Println(fn(5))
	fmt.Println(fn(timeToWait))
}
