package main

import "fmt"

var num int

func main() {
	cards := newDeck()

	hand, cards := deal(cards, 7)

	hand.print()
	fmt.Println("---------------")
	cards.print()

	fmt.Println("---------------")

	fmt.Println(hand.toString())
}
