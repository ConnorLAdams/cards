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
	fmt.Println("")

	player1 := players.Player{
		Name:  "Goseph",
		Cards: newDeck.DealToPlayer(),
	}

	player1.ShowCards()

	player2 := players.Player{
		Name: "Rusty",
		Cards: newDeck.DealToPlayer(),
	}

	player2.ShowCards()

	fmt.Println(compareHands(player1, player2))
}

func compareHands(p1 players.Player, p2 players.Player) string {
	p1Cards := p1.Cards
	p2Cards := p2.Cards
	
	fmt.Println(p1Cards, p2Cards)
	fmt.Println("")
	return p1.Name
}
