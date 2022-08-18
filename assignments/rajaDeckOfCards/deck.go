package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Card struct {
	numbers string
	suit    string
}
type Deck []Card

func createNewDeck() (deck Deck) {
	suits := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	// numbers := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	numbers := []string{"Ace", "Two", "Three", "Four"}

	// for i := 0; i < len(numbers); i++ {
	// 	for j := 0; j < len(suits); j++ {
	// 		card := numbers[i] + " of " + suits[j]
	// 		deck = append(deck, card)
	// 	}
	// }

	// index, value
	for _, num := range numbers {
		for j := range suits {
			card := Card{
				numbers: num,
				suit:    suits[j],
			}
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

	// how to switch / swap values?
	// d[0], d[15] = d[15], d[0]

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range d {
		// find a random value for the new position between 0 and len - 1
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

// return deck of cards distributed and remaining
func deal(d Deck, handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

func (c Card) toString() string {

	return c.numbers + " of " + c.suit

}

func (d Deck) saveToFile(fileName string) error {
	// combine all the elements to single string comma separated
	var str string

	for _, c := range d {
		str = c.toString() + "," + str
	}

	return os.WriteFile("myDeck.txt", []byte(str), 0666)
}

// func(d Deck) toCard(s string) Card {}
func ReadFromFile(fileName string) Deck {
	data, err := os.ReadFile("myDeck.txt")

	if err != nil {
		fmt.Println("Error while saving to file", err)
		os.Exit(1)
	}
	initialDeck := strings.Split(string(data), ",")
	var finalDeck = make(Deck, 0)
	for _, i := range initialDeck {
		inside := strings.Split(i, " of ")
		if len(inside) == 2 {
			finalDeck = append(finalDeck, Card{inside[0], inside[1]})
		}
	}
	return finalDeck
}
