package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected length of the deck is 52 cards, not %v", len(d))
	}

}

func TestFirstCard(t *testing.T) {
	d := newDeck()
	first_card := "Ace of Diamonds"
	if d[0] != first_card {
		t.Errorf("Expected first card %v, not %v", first_card, d[0])
	}
}

func TestLastCard(t *testing.T) {
	d := newDeck()
	last_card := "King of Spades"
	if d[len(d)-1] != last_card {
		t.Errorf("Expected last card %v, not %v", last_card, d[len(d)-1])
	}
}

func TestDeckOperations(t *testing.T) {
	//deleting test file if it somehow exists
	os.Remove("_decktesting")

	d := newDeck()

	d.saveToFile("_decktesting")

	b := readDeck("_decktesting")

	if len(b) != 52 {
		t.Errorf("Incorrect amount of cards in deck, expected 52, not %v", len(b))
	}

	os.Remove("_decktesting")
}
