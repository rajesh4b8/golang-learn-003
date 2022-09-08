package main

import (
	"fmt"
	"time"
)

func main() {
	demoSelect()
}

func demoSwitch() {
	for i := 0; i < 5; i++ {
		switch i {
		case 1:
			fmt.Println("It's one")
		case 2:
			fmt.Println("It's two")
		case 3:
			fmt.Println("It's three")
		default:
			fmt.Println("Default!")
		}
	}

	// Empty interface.. everything can justify the implementation
	var a interface{}
	a = true

	// type checking
	switch a.(type) {
	case int:
		fmt.Println("It's an int")
	case bool:
		fmt.Println("It's a boolean")
	default:
		fmt.Println("Don't know")
	}
}

func demoSelect() {
	tick := time.Tick(1000 * time.Millisecond)
	boom := time.After(5000 * time.Millisecond)

	// infinte loop
	for {
		select {
		case <-tick:
			fmt.Println("Tick")
		case <-boom:
			fmt.Println("Boom")
			// exit from loop
			return
		default:
			fmt.Println("   .")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
