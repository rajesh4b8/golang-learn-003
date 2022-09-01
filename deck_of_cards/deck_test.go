package main

import "testing"

func TestNewDeck(t *testing.T) {
	newDeck := createNewDeck()

	if len(newDeck) != 16 {
		t.Errorf("Expected 16 cards but got %d", len(newDeck))
	}

	if newDeck[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades but got %v", newDeck[0])
	}

	if newDeck[15] != "Four of Hearts" {
		t.Errorf("Expected Four of Hearts but got %v", newDeck[15])
	}
}

func TestDeal(t *testing.T) {
	// Create a new deck
	// call the deal with size
	// test if the deal size is expected
	// test if the remaing cards size is expected
}
