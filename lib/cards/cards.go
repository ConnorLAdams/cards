package cards

import (
	"math/rand"
	"time"
	"fmt"
)

type Hand [2]string
type Deck []string
type Cards interface {
	ShowCards()
}

func (d Deck) Shuffle() {
	seed := rand.NewSource(time.Now().UnixNano())
	Rand := rand.New(seed)
	for i := len(d) - 1; i > 0; i-- {
		j := Rand.Intn((i + 1))
		d[i], d[j] = d[j], d[i]
	}
}

func (d Deck) DealToPlayer() Hand {
	seed := rand.NewSource(time.Now().UnixNano())
	Rand := rand.New(seed)

	var newHand Hand

	card1Index := Rand.Intn(len(d))
	card1 := d[card1Index]
	newHand[0] = card1
	d = append(d[:card1Index], d[card1Index+1:]...)

	card2Index := Rand.Intn(len(d))
	card2 := d[card2Index]
	newHand[1] = card2
	d = append(d[:card2Index], d[card2Index+1:]...)

	return newHand
}

func (h Hand) ShowCards() {
	fmt.Println(h)
}

func (d Deck) ShowCards() {
	fmt.Println(d)
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
