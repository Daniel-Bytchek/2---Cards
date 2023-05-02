package main

var num int

func main() {
	cards := newDeck()

	hand, cards := deal(cards, 7)

	hand.saveToFile("player1")
}
