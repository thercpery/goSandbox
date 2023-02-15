package main

import "fmt"

type ContactInfo struct {
	email   string
	zipCode int
}

type Person struct {
	firstName string
	lastName  string
	ContactInfo
}

func main() {
	alex := Person{
		firstName: "Alex",
		lastName:  "Anderson"}
	fmt.Println(alex)

	var bob Person
	bob.firstName = "Bob"
	bob.lastName = "Smith"
	fmt.Println(bob)
	fmt.Printf("%+v\n", bob)

	jim := Person{
		firstName: "Jim",
		lastName:  "Party",
		ContactInfo: ContactInfo{
			email:   "jim@gmail.com",
			zipCode: 87000,
		},
	}

	// jimPointer := &jim // "&" gives the RAM address to "jim" variable
	jim.updateName("Jimbo")
	jim.print()

	/*
		&variable = gives the memory address of the value the variable is pointing at
		*pointer = gives the value of the memory address is pointing at

		pointer = memory address
		value = value
	*/
}

func (pointerToPerson *Person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p Person) print() {
	fmt.Printf("%+v\n", p)
}
