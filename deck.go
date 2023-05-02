package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func readDeck(fileName string) deck {
	stringDeck, error := ioutil.ReadFile(fileName)
	if error != nil {
		fmt.Println("Error happened! ", error)
		os.Exit(1)
	}
	stringSlice := strings.Split(string(stringDeck), "^")

	return deck(stringSlice)
}
