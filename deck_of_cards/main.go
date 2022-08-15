package main

import (
	"fmt"
	"os"
)

func main() {
	deck := createNewDeck()
	deck.shuffle()
	// deck.print()

	// deal, remaining := deal(deck, 5)
	// fmt.Println("Deal pack")
	// deal.print()
	// fmt.Println("Remaining pack")
	// remaining.print()

	err := deck.saveToFile("myDeck.txt")
	if err != nil {
		fmt.Println("Error while saving to file", err)
		os.Exit(1)
	}

	d1 := ReadFromFile("myDeck.txt")
	d1.print()
}
