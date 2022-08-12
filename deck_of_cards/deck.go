package main

import "fmt"

type Deck []string

func createNewDeck() Deck {
	var deck []string
	suits := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	numbers := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// TODO create a deck of cards and return
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(suits); j++ {
			card := numbers[i] + " of " + suits[j]
			deck = append(deck, card)
		}
	}
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
