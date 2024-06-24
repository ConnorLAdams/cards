package cards

import (
	"math/rand"
	"time"
)

type Hand [2]string
type Deck []string

func (d Deck) Shuffle() {
	seed := rand.NewSource(time.Now().UnixNano())
	Rand := rand.New(seed)
	for i := len(d) - 1; i > 0; i-- {
		j := Rand.Intn((i + 1))
		d[i], d[j] = d[j], d[i]
	}
}

func GenerateDeck() Deck {
	suites := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	values := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}

	var cards Deck
	for _, suit := range suites {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}
