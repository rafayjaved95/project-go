package main

func main() {
	// cards := newDeck()

	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// remainingDeck.print()

	// cards := newDeck()
	// fmt.Println(cards.toString())

	// cards.saveToFile("my_cards")

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
