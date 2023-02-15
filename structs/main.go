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

	jim.print()
	jim.updateName("Jimbo")
}

func (p Person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p Person) print() {
	fmt.Printf("%+v\n", p)
}
