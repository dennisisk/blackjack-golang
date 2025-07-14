package main

import (
	"blackjack/game"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	balance := 100
	fmt.Println("\n🎲 Welcome to Go Blackjack!")

	for balance > 0 {
		fmt.Printf("\n💰 Current balance: $%d\n", balance)
		bet := game.PromptBet(balance)

		deck := game.NewDeck()
		playerHand := []game.Card{deck.Draw(), deck.Draw()}
		dealerHand := []game.Card{deck.Draw(), deck.Draw()}

		fmt.Println("\nYour hand:")
		game.DisplayHand(playerHand)
		fmt.Printf("Total: %d\n", game.CalculateScore(playerHand))

		fmt.Println("\nDealer shows:")
		game.DisplayCard(dealerHand[0])

		// Player turn
		for game.CalculateScore(playerHand) < 21 {
			choice := game.PromptChoice()
			if choice == "h" {
				playerHand = append(playerHand, deck.Draw())
				fmt.Println("\nYou drew:")
				game.DisplayCard(playerHand[len(playerHand)-1])
				fmt.Printf("Total: %d\n", game.CalculateScore(playerHand))
			} else {
				break
			}
		}

		playerScore := game.CalculateScore(playerHand)
		if playerScore > 21 {
			fmt.Println("\n💥 Bust! You lose this round.")
			balance -= bet
		} else {
			// Dealer turn
			fmt.Println("\nDealer's hand:")
			game.DisplayHand(dealerHand)
			for game.CalculateScore(dealerHand) < 17 {
				card := deck.Draw()
				dealerHand = append(dealerHand, card)
				fmt.Println("Dealer draws:")
				game.DisplayCard(card)
			}
			dealerScore := game.CalculateScore(dealerHand)
			fmt.Printf("Dealer final: %d\n", dealerScore)

			switch {
			case dealerScore > 21 || playerScore > dealerScore:
				fmt.Println("\n✅ You win!")
				balance += bet
			case playerScore < dealerScore:
				fmt.Println("\n❌ Dealer wins.")
				balance -= bet
			default:
				fmt.Println("\n⚖️  It's a tie.")
			}
		}

		if balance <= 0 {
			fmt.Println("\n😢 You're out of money!")
			break
		}

		fmt.Print("\n🔁 Play another round? (y/n): ")
		var again string
		fmt.Scanln(&again)
		if again != "y" {
			break
		}
	}
	fmt.Println("\n👋 Thanks for playing Go Blackjack!")
}
