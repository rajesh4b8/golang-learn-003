package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p := person{"Rajesh", 12}

	fmt.Printf("p:: %+v\n", p)

	personPointer := &p
	fmt.Printf("personPointer:: %p\n", personPointer)

	valueAtPointer := *personPointer
	fmt.Printf("valueAtPointer:: %+v\n", valueAtPointer)

	p.updateNameByValue()
	fmt.Printf("p after updateNameByValue:: %+v\n", p)

	// (&p).updateNameByPointer()
	p.updateNameByPointer() // pointer indirection -> the pointer to the person would be passed automatically
	fmt.Printf("p after updateNameByPointer:: %+v\n", p)
}

// pass by value -> data is copied
func (p person) updateNameByValue() {
	fmt.Printf("personPointer inside updateNameByValue():: %p\n", &p)
	p.name = "Suresh"
}

// pass by pointer -> pointer referes to same data
func (pp *person) updateNameByPointer() {
	fmt.Printf("personPointer inside updateNameByPointer():: %p\n", &pp)
	// (*pp).name = "Suresh"
	pp.name = "Suresh" // pointer indirection
}
