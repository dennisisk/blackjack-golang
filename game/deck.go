package game

import (
	"math/rand"
)

type Card struct {
	Suit  string
	Value string
}

type Deck []Card

func NewDeck() *Deck {
	suits := []string{"♠", "♥", "♦", "♣"}
	values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	deck := Deck{}
	for _, suit := range suits {
		for _, val := range values {
			deck = append(deck, Card{Suit: suit, Value: val})
		}
	}
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})

	return &deck
}

func (d *Deck) Draw() Card {
	card := (*d)[0]
	*d = (*d)[1:]
	return card
}
