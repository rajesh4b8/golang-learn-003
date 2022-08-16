package main

import "fmt"

type person struct {
	name string
}

func main() {
	// the zero value for int will be 0
	var j int
	p := &j

	// same as above but more convenient
	// new() -> initialize with zero value and return the pointer
	i := new(int)

	fmt.Println("i", i)
	fmt.Println("p", p)
	fmt.Println(*i)
	fmt.Println(*p)

	// var m1 map[string]string -> zero value is nil
	// mp1 := &m1
	mp1 := new(map[string]string)
	fmt.Println("mp1", mp1)
	fmt.Println("*mp1", *mp1)

	// it's panic when you try update a map created using new keyword
	// (*mp1)["test"] = "Something"

	// always use make() to create maps
	mp2 := make(map[string]string)
	mp2["test"] = "Something"
	fmt.Println("mp2", mp2)

	// var person p
	// p1 := &p
	p1 := new(person)
	// (*p1).name = "Rajesh"
	p1.name = "Rajesh" // pointer indirection
	fmt.Println(p1)
	fmt.Println(*p1)
}
