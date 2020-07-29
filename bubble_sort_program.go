// Write a Bubble Sort program in Go. The program should prompt the user to type in a sequence of up to 10 integers. The program should print the integers out on one line, in sorted order, from least to greatest. Use your favorite search tool to find a description of how the bubble sort algorithm works.

// As part of this program, you should write a function called BubbleSort() which takes a slice of integers as an argument and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted order.

// A recurring operation in the bubble sort algorithm is the Swap operation which swaps the position of two adjacent elements in the slice. You should write a Swap() function which performs this operation. Your Swap() function should take two arguments, a slice of integers and an index value i which indicates a position in the slice. The Swap() function should return nothing, but it should swap the contents of the slice in position i with the contents in position i+1.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hi there!")

	sliceLen := 10
	slice := make([]int64, 0, 10)

	fmt.Println("Type ten integers to sort...")

	var num string
	for m := 0; m < sliceLen; m++ {
		fmt.Scan(&num)
		num, err := strconv.ParseInt(num, 10, 64)

		if err != nil {
			fmt.Println("I only accept integers, try again...")
			return
		}

		slice = append(slice, num)
	}

	BubbleSort(slice)
	fmt.Println(slice)
}

// BubbleSort provided slice
func BubbleSort(slice []int64) {
	size := len(slice)

	for index := 1; index < size; index++ {
		switcher := false

		for i := 0; i < size-index; i++ {
			nextIndex := i + 1
			if slice[i] > slice[nextIndex] {
				Swap(slice, i)
				switcher = true
			}
		}

		if switcher == false {
			break
		}
	}
}

// Swap values in provided slice
func Swap(ints []int64, i int) {
	ints[i], ints[i+1] = ints[i+1], ints[i]
}
