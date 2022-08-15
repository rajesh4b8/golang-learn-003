package main

import (
	"fmt"
)

// definition of struct
type person struct {
	firstName string
	lastName  string
	age       int
	addr      address
}

type address struct {
	city string
	zip  int
}

func main() {
	// declaration of struct (person) without initialization
	// properties inside would initialized with zero values
	var p person

	// var p1 person
	p1 := person{}

	fmt.Println(p, p1)

	// p2 := person{firstName: "Rajesh", lastName: "Reddy", age: 12}
	// p2 := person{"Rajesh", "Reddy", 12}
	// p2 := person{
	// 	"Rajesh",
	// 	"Reddy",
	// 	12,
	// }

	p2 := person{
		"Rajesh",
		"Reddy",
		12,
		address{"Phoeniz", 85021},
	}

	fmt.Printf("p2:: %+v\n", p2)

	// p3 := person{"", "Test", 0}
	p3 := person{lastName: "Test"}

	p3.firstName = "abc"
	p3.addr.city = "Dallas"
	fmt.Printf("p3:: %+v\n", p3)

	// updatePersonFunc(p3)
	p3.updatePerson()
	fmt.Printf("p3 after calling updatePerson():: %+v\n", p3)
}

// pass by value -> p3 is copied to p
func updatePersonFunc(p person) {
	p.firstName = "updated!"
	fmt.Printf("p3 inside updatePerson():: %+v\n", p)
}

// pass by value -> p3 is copied to p
func (p person) updatePerson() {
	p.firstName = "updated!"
	fmt.Printf("p3 inside updatePerson():: %+v\n", p)
}
