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
	contactInfo
	// or contact contactInfo
}

// 2. create a new value of the struct type
func main() {
	harry := person{firstName: "Harry",
		lastName: "Potter",
		contactInfo: contactInfo{
			email:    "harry@gmail.com",
			postcode: 358000,
		},
	}
	// this points to the actual address of the harry person so that it gets properly updated
	harryPointer := &harry
	harryPointer.updateName("Ron")

	harry.printInfo()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) printInfo() {
	fmt.Printf("%+v", p)
}

// as soon as you create a struct person, Go will go and try to find some space on your local
// machine to store this data in a container.

// Go is a pass by value language. Whenever we pass some value into a func, Go will take it
// and copy all of this updated data into a NEW empty container.

// & - give me actual memory address of the value of this struct/variable
// * - give me the actual value of this memory that address is pointing at
