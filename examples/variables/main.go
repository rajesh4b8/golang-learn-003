package main

import "fmt"

var (
	t = 12
	a = 50
)

func main() {
	// var is the keyword
	// foo is the name of the varalbe
	// int is built in type
	var foo int = 42    // declaring variable "foo" as int and initializing with 42
	var foo1 = 42       // declaration with initialization -> type is determined by data
	var bar int         // declaration without initialization
	var i, j int = 2, 4 // multiple values
	var x, y int

	// changing the variable data
	foo = 45

	// var name string = "Rajesh"
	name := "Rajesh" // this exactly same as line above... shortcut

	// changing the name and name is already declared
	name = "Ramesh"

	fmt.Println(foo, foo1, bar, name, i, j, x, y)
	fmt.Println("main()", a)
	a = 12
	fmt.Println("main()", a)

	printSomething()
}

func printSomething() {
	a := 12
	fmt.Println("printSomething()", a)
}
