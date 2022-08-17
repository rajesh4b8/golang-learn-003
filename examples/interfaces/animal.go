package main

import (
	"fmt"
)

// Animal has a Name and an Age to represent an animal.
// It implements the Stringer interface by implementing the method `String() string`
type Animal struct {
	Name string
	Age  uint
}

// String makes Animal satisfy the Stringer interface.
func (a Animal) String() string {
	return fmt.Sprintf("Implmentation of String: %v (%d)", a.Name, a.Age)
}

// func main() {
func animalTest() {
	a := Animal{
		Name: "Gopher",
		Age:  2,
	}

	fmt.Println(a)

	printAnything(a)
}

func printAnything(s fmt.Stringer) {
	fmt.Println("inside printAnything", s.String())
}
