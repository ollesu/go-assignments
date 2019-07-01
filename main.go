package main

import "fmt"

func main() {
	// this initialises a slice of length of 11
	ints := make([]int, 11)

	// making a slice
	for i := 0; i < 11; i++ {
		ints[i] = i
	}

	for _, num := range ints {
		if num%2 == 0 {
			fmt.Println(num, "is even")
		} else {
			fmt.Println(num, "is odd")
		}
	}
}
