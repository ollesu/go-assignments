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
	// not referencing directly to the struct, but the memory address
	// harryPointer := &harry
	// harryPointer.updateName("Ron")

	// shortcut as Go can automatically turn var of person type into pointer of person:
	harry.updateName("Ron")
	harry.printInfo()
}

// as soon as you create a struct person, Go will go and try to find some space on your local
// machine to store this data in a container.

// Go is a pass by value language. Whenever we pass some value into a func, Go will take it
// and copy all of this updated data into a NEW empty container.

// &variable - give me actual memory address of the value of this struct/variable
// *pointer - give me the actual value of this memory that address is pointing at

// *person - this a description of a data type - i.e. a type of person pointer
// this means updateName func can only be called with a type of person pointer
func (pointerToSomething *person) updateName(newFirstName string) {
	(*pointerToSomething).firstName = newFirstName
}

func (p person) printInfo() {
	fmt.Printf("%+v", p)
}
