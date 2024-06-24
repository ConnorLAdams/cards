package players

import (
	cards "cards/lib/cards"
	"fmt"
)

type Player struct {
	Name  string
	Cards cards.Hand
}

func (p Player) ShowCards() {
	cards_str := p.Cards[0] + ", " + p.Cards[1]
	fmt.Println(p.Name + " Holds the following cards: " + cards_str)
}
