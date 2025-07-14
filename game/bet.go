package game

import (
	"fmt"
)

func PromptBet(balance int) int {
	var bet int
	for {
		fmt.Print("Enter your bet: $")
		fmt.Scanln(&bet)
		if bet > 0 && bet <= balance {
			return bet
		}
		fmt.Println("Invalid bet amount.")
	}
}

func PromptChoice() string {
	var choice string
	for {
		fmt.Print("Hit or Stand? (h/s): ")
		fmt.Scanln(&choice)
		if choice == "h" || choice == "s" {
			return choice
		}
		fmt.Println("Invalid choice.")
	}
}
