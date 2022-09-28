package main

func main() {
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	// hand, remaininCards := deal(cards, 5)
	// cards.saveToFile("my_cards")
	cards.print()
}
