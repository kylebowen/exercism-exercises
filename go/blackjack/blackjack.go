package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	playerHandValue := ParseCard(card1) + ParseCard(card2)
	dealerCardValue := ParseCard(dealerCard)

	switch {
	case playerHandValue > 21:
		return "P"
	case playerHandValue == 21 && dealerCardValue < 10:
		return "W"
	case playerHandValue == 21:
		return "S"
	case 17 <= playerHandValue && playerHandValue <= 20:
		return "S"
	case 12 <= playerHandValue && playerHandValue <= 16 && dealerCardValue >= 7:
		return "H"
	case 12 <= playerHandValue && playerHandValue <= 16:
		return "S"
	case playerHandValue <= 11:
		return "H"
	default:
		return "How did you even?"
	}
}
