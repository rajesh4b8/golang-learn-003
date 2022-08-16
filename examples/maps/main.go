package main

import "fmt"

func main() {

	// Map declaration -> the zero value is nil -> don't use this way
	// var shapes map[string]int

	// Always use make() to create a map
	shapes := make(map[string]int)

	shapes["triangle"] = 3

	fmt.Println(shapes)

	colors := map[string]string{
		"Red":  "#d463424",
		"Blue": "#64322",
	}
	colors["Red"] = "#63432"
	updateColor(colors)

	printMap(colors)

	// compile issue if the type is not same
	// printMap(shapes)

}

func updateColor(c map[string]string) {
	c["Black"] = "#000000"
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
