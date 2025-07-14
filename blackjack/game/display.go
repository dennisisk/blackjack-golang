package game

import (
	"fmt"
)

func DisplayHand(hand []Card) {
	for i := 0; i < 5; i++ {
		for _, card := range hand {
			switch i {
			case 0:
				fmt.Print(" ____ ")
			case 1:
				fmt.Printf("|%-2s  |", card.Value)
			case 2:
				fmt.Printf("| %s  |", card.Suit)
			case 3:
				fmt.Printf("|  %-2s|", card.Value)
			case 4:
				fmt.Print("|____|")
			}
		}
		fmt.Println()
	}
}

func DisplayCard(card Card) {
	hand := []Card{card}
	DisplayHand(hand)
}
