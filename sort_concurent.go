package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hi there!")

	sliceLen := 10
	slice := make([]int, 0, 10)

	fmt.Println("Type ten integers to sort...")

	var num string
	for m := 0; m < sliceLen; m++ {
		fmt.Scan(&num)
		num, err := strconv.ParseInt(num, 10, 64)

		if err != nil {
			fmt.Println("I only accept integers, try again...")
			return
		}

		slice = append(slice, int(num))
	}

	o := 0
	size := len(slice)
	quarter := size / 4

	// Split slice into four parts
	temp := make([]int, 0, size)
	result := make([]int, 0, size)

	for i := 0; i < 4; i++ {
		firstChannel := make(chan []int)

		if i == 3 {
			temp = slice[o:size]
		} else {
			temp = slice[o:(o + quarter)]
		}

		go BubbleSort(temp, firstChannel)
		for res := range firstChannel {
			result = append(result, append(res)...)
		}
		o += quarter
	}

	finalChannel := make(chan []int)
	go BubbleSort(result, finalChannel)
	fmt.Println("Now performing the final sort")
	fmt.Println(<-finalChannel)
}

// BubbleSort provided slice
func BubbleSort(slice []int, channel chan []int) {
	defer close(channel)

	fmt.Printf("This will sort the following chunk: %v\n", slice)

	size := len(slice)

	for index := 1; index < size; index++ {
		switcher := false

		for i := 0; i < size-index; i++ {
			nextIndex := i + 1
			if slice[i] > slice[nextIndex] {
				slice[i], slice[i+1] = slice[i+1], slice[i]
				switcher = true
			}
		}

		if switcher == false {
			break
		}
	}

	channel <- slice
}
