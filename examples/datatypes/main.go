package main

import "fmt"

func main() {

	var flag bool = true
	var str string = "Something text"
	var i int = -1234
	var ui uint = 1234
	var r rune = 1234 // =~ int32
	var b byte = 255  // byte =~ uint8
	var f float32 = 2345.356

	fmt.Println(flag, str, i, ui, r, b, f)
}
