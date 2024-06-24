package players

import (
	cards "cards/lib/cards"
	"fmt"
)

type player struct {
	name  string
	cards cards.Hand
}

func (p player) ShowCards() {
	fmt.Println(p.cards)
}
