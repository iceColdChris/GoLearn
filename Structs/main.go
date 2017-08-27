package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string

	//Struct embedding
	contact contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func main() {
	//This depends on order
	//alex := person{"Alex", "Anderson"}

	//Field assignment
	//alex := person{firstName: "Alex", lastName: "Anderson"}

	//Pass by field
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	//Creating an embedded struct
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email: "jim@gmail.com",
			zip:   94000,
		},
	}

	fmt.Println(jim)

	//Prints out field names and values
	//fmt.Printf("%+v\n", jim)

	jim.print()
}

func (p person) updateName(firstName string) {
	p.firstName = firstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
