package main

import (
	cards "cards/lib/cards"
	players "cards/lib/players"
	"fmt"
)

func main() {
	fmt.Println("Let's play some cards")

	newDeck := cards.GenerateDeck()
	newDeck.Shuffle()
	fmt.Println(newDeck)

	newPlayer := players.Player{
		Name:  "Goseph",
		Cards: newDeck.DealToPlayer(),
	}

	newPlayer.ShowCards()
}
