package game

import "strconv"

func CalculateScore(hand []Card) int {
	score := 0
	aces := 0
	for _, card := range hand {
		switch card.Value {
		case "A":
			score += 11
			aces++
		case "K", "Q", "J":
			score += 10
		default:
			val, _ := strconv.Atoi(card.Value)
			score += val
		}
	}
	for score > 21 && aces > 0 {
		score -= 10
		aces--
	}
	return score
}
