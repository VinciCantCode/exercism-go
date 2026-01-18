package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	result := 0
	switch {
	case card == "ace":
		result = 11
	case card == "two":
		result = 2
	case card == "three":
		result = 3
	case card == "four":
		result = 4
	case card == "five":
		result = 5
	case card == "six":
		result = 6
	case card == "seven":
		result = 7
	case card == "eight":
		result = 8
	case card == "nine":
		result = 9
	case card == "ten":
		result = 10
	case card == "jack":
		result = 10
	case card == "queen":
		result = 10
	case card == "king":
		result = 10
	case card == "other":
		result = 0
	}
	return result
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var result string
	switch {
	case card1 == "ace" && card2 == "ace":
		result = "P"
	case ParseCard(card1)+ParseCard(card2) == 21 && ParseCard(dealerCard) != 10 && ParseCard(dealerCard) != 11:
		result = "W"
	case (ParseCard(card1)+ParseCard(card2) == 21) && (ParseCard(dealerCard) == 10 || ParseCard(dealerCard) == 11):
		result = "S"
	case ParseCard(card1)+ParseCard(card2) >= 17 && ParseCard(card1)+ParseCard(card2) <= 20:
		result = "S"
	case ParseCard(card1)+ParseCard(card2) >= 12 && ParseCard(card1)+ParseCard(card2) <= 16 && ParseCard(dealerCard) < 7:
		result = "S"
	case ParseCard(card1)+ParseCard(card2) >= 12 && ParseCard(card1)+ParseCard(card2) <= 16 && ParseCard(dealerCard) >= 7:
		result = "H"
	case ParseCard(card1)+ParseCard(card2) <= 11:
		result = "H"
	}
	return result
}
