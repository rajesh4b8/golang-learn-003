package main

import (
	"fmt"
	"unsafe"
)

type emty struct{}

type person struct {
	firstName string
	lastName  string
}

func main() {
	var e emty
	f := emty{}
	g := struct{}{}

	fmt.Println(e, f, g)
	fmt.Println("Size of e", unsafe.Sizeof(e))
	fmt.Println("e == f", e == f)
	fmt.Println("g == f", g == f)

	p1 := person{"Rajesh", "Reddy"}
	fmt.Println("p1", p1)
	fmt.Println("Size of p1", unsafe.Sizeof(p1))

	b1 := true
	fmt.Println("Size of b1", unsafe.Sizeof(b1))

}
