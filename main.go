package main

var num int

func main() {
	deck := newDeck()
	deck.shuffle()
	deck.print()
}
