package main

import (
	"fmt"
	"time"
)

func main() {

	// student's note - this is an example of go routine race condition.
	// If run on a fast device you may happen to see both statements printed
	// But most of the time only one will be visible. This demonstrates unpredictability.

	go fmt.Println("How are you doing?")
	time.Sleep(1)
	fmt.Println("Hello there!")
}
