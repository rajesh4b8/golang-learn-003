package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 42
	var f float32
	f = float32(i)

	k := uint(i)

	fmt.Println(f, k)

	x := "7645"
	y, err := strconv.Atoi(x)
	if err != nil {
		panic(err)
	}
	fmt.Println(y)

	str := "test"

	fmt.Println("Byte Slice:", []byte(str))
}
