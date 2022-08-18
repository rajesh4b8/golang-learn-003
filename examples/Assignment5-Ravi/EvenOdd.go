//Print numbers from 1 to 10 and print whether it's even or odd
package main

import (
	"fmt"
)

func main() {
	for num := 0; num < 11; num++ {
		if num%2 == 0 {
			fmt.Printf("%v is Even\n", num)
		} else {
			fmt.Printf("%v is odd\n", num)

		}
	}

}
