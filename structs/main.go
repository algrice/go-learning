package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	amanda := person{
		firstName: "Amanda",
		lastName:  "Grice",
		contact: contactInfo{
			email:   "fake@gmail.com",
			zipCode: 12345,
		},
	}
	amanda.print()
	amanda.updateName("Mandy")
	amanda.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	pointerToPerson.firstName = newFirstName
}
