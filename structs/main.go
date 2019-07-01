package main

import "fmt"

// 1. define properties of a struct
type contactInfo struct {
	email    string
	postcode int
}

// here's a first embedded struct
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

// 2. create a new value of the struct type
func main() {
	harry := person{firstName: "Harry",
		lastName: "Potter",
		contact: contactInfo{
			email:    "harry@gmail.com",
			postcode: 358000,
		},
	}
	fmt.Printf("%+v", harry)
}
