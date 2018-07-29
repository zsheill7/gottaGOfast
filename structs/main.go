package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode string
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

// initially assigns zero value

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Halpert",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: "98040",
		},
	}

	jim.updateName("jimmy")
	jim.print()
}

// The ampersand says give me the address this pointer is pointing to
// * Give the value that the address is pointing to

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
