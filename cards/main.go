package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remainingDeck := deal(cards, 5)

	fmt.Println("\nHand:")
	hand.print()

	fmt.Println("\nRemaining Deck:")
	remainingDeck.print()

	fmt.Println(cards.toString())
	cards.saveToFile("my_cards.txt")

	deck := newDeckFromFile("my_cards.txt")
	deck.print()
	fmt.Println()
	deck.shuffle()
	deck.print()
}
