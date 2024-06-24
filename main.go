package main

import (
	cards "cards/lib/cards"
	"fmt"
)

func main() {
	fmt.Println("Let's play some cards")

	newDeck := cards.GenerateDeck()
	newDeck.Shuffle()
	fmt.Println(newDeck)
}
