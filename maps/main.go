package main

import "fmt"

func main() {
	// CREATE:
	// 1. colors := make(map[string]string)
	// 2. [string] refers to keys, string refers to values:
	colours := map[string]string{
		"tulip":        "#ff0000",
		"darkorchid":   "#9932cc",
		"spanish grey": "94a89a",
	}
	// UPDATE
	colours["desert sand"] = "#edc9af"

	// DELETE
	delete(colours, "tulip")

	printMap(colours)
}

func printMap(c map[string]string) {
	for colour, hexCode := range c {
		fmt.Println("Hex code for", colour, "is", hexCode)
	}
}
