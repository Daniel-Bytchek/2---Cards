package main

import "testing"

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
