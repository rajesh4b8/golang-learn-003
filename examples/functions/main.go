package main

import "fmt"

func main() {
	simpleFunction()

	withParams("test", 23)

	str := withReturnSingle()
	fmt.Println("Returned from withReturnSingle: ", str)

	// var name string
	// var age int
	// name, age = withReturnMutltiple()
	name, age := withReturnMutltiple()

	fmt.Printf("withReturnMutltiple name: %s and age: %d\n", name, age)

	fmt.Println("2 + 3 =", add(2, 3))
	fmt.Println("2 + 3 =", applyOperation(add, 2, 3))

	substract := func(i int, j int) int {
		fmt.Println("inside subtract", i, j)
		return i - j
	}

	k := substract(2, 3)
	fmt.Println("2 - 3 =", k)
	fmt.Println("2 - 3 =", applyOperation(substract, 2, 3))

	// inline functions calling immediately
	fmt.Println("2 * 3 =", func(i int, j int) int {
		return i * j
	}(2, 3))
}

func simpleFunction() {
	fmt.Println("simpleFunction")
}

func withParams(p1 string, p2 int) {
	fmt.Printf("withParams p1: %s and p2: %d\n", p1, p2)
}

func withReturnSingle() string {
	return "something"
}

func withReturnMutltiple() (string, int) {
	return "Rajesh", 12
}

func add(x int, y int) int {
	return x + y
}

func applyOperation(operation func(i, j int) int, x int, y int) int {
	return operation(x, y)
}
