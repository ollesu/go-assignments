package main

import "fmt"

func main() {
	// this initialises a slice of length of 11
	ints := make([]int, 11)

	// making a slice
	for i := 0; i < 11; i++ {
		ints[i] = i
	}
	fmt.Println(ints)
}
