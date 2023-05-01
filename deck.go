package main

import (
	"fmt"
	"strings"
)

type deck []string

func newDeck() deck {
	newCards := deck{}

	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	cardSuits := []string{"Diamonds", "Hearts", "Clubs", "Spades"}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			newCards = append(newCards, cardValue+" of "+cardSuit)
		}
	}

	return newCards
}

func deal(d deck, amount int) (deck, deck) {
	return d[:amount], d[amount:]
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), "^")
}
