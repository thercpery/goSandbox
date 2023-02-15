package main

import "fmt"

type ContactInfo struct {
	email   string
	zipCode int
}

type Person struct {
	firstName string
	lastName  string
	contact   ContactInfo
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
	fmt.Printf("%+v", bob)

	jim := Person{
		firstName: "Jim",
		lastName:  "Party",
		contact: ContactInfo{
			email:   "jim@gmail.com",
			zipCode: 87000,
		},
	}

	fmt.Printf("%+v", jim)
}
