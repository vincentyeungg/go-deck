package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)

	fmt.Println("Printing hand:")
	hand.print()

	// fmt.Println("Printing remaining deck:")
	remainingDeck.print()

	fmt.Println(cards.toString())

	cards.saveToFile("my_cards")
}
