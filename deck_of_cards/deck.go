package main

import "fmt"

type Deck []string

func createNewDeck() Deck {
	card1 := "Ace of Spades"
	card2 := "Two of Diamonds"

	var deck []string = []string{card1, card2}
	deck = append(deck, "Four of Hearts")

	// suits := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	// numbers := []string{"Ace", "Two", "Three", "Four"}

	// TODO create a deck of cards and return

	return deck
}

// this is function and deck passed as param
func printFunc(d Deck) {
	// blank reference `_` if you don't want to use index
	for _, card := range d {
		fmt.Println(card)
	}
}

// this is method defined on type "Deck"
func (d Deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d Deck) shuffle() {
	// iterate over all the elements
	// switch with a random position

	// how to switch values?
	// d[0], d[15] = d[15], d[0]
}
