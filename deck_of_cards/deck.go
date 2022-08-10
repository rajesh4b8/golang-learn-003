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

func print(d Deck) {
	for i, card := range d {
		fmt.Println(i, card)
	}

	// blank reference `_` if you don't want to use index
	for _, card := range d {
		fmt.Println(card)
	}
}
