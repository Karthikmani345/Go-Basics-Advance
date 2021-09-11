package main

import (
	"fmt"
)

func main() {
	// var emp1 Person
	// emp1.firstName = "karthik"
	// emp1.lastName = "Mani"

	emp1 := NewPerson("kartik", "mani")

	fmt.Println(emp1)
	fmt.Println(emp1.GetFullname())
	emp1.updateName("murthi")
	fmt.Println(emp1)
	fmt.Println(emp1.GetFullname())
}

type Person struct {
	firstName string
	lastName  string
}

// constructor like in object oriented programming
func NewPerson(firstName string, lastName string) Person {
	return Person{firstName, lastName}
}

/// X wont work
// func (p Person) New(firstName string, lastName string) Person {
// 	return Person{firstName, lastName}
// }

// for readonly i.e getter
func (p Person) GetFullname() string {
	return p.firstName + " " + p.lastName
}

// for setting the value need to pass the pointer i.e setter
func (p *Person) updateName(value string) {
	p.firstName = value
}
