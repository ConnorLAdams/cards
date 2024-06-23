package cards

import (
	"math/rand"
	"time"
)

// type hand [2]string
type deck []string

func (d deck) Shuffle() {
	seed := rand.NewSource(time.Now().UnixNano())
	Rand := rand.New(seed)
	for i := len(d) - 1; i > 0; i-- {
		j := Rand.Intn((i + 1))
		d[i], d[j] = d[j], d[i]
	}
}

func GenerateDeck() deck {
	suites := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	values := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}

	var cards deck
	for _, suit := range suites {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}
