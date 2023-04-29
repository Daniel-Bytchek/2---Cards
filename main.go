package main

import (
	"fmt"
	"strconv"
)

var num int

func main() {
	num = 0
	cards := []string{"Ace of Spades", newCard()}
	cards = append(cards, "Ace of Diamonds")

	for i, card := range cards {
		i = i
		fmt.Println(setNum(card))
	}
}

func newCard() string {
	return "Five of Diamonds"
}

func setNum(card string) string {
	num++
	return strconv.Itoa(num) + " " + card
}
