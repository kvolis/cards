package cards

// NewDeck36 returns card's deck from Six to Ace by all suits
func NewDeck36() Cards {
	cards, i := make(Cards, 36), 0
	for rank := Six; rank <= Ace; rank++ {
		for suit := Hearts; suit <= Clubs; suit += 1 << 4 {
			cards[i] = Card(rank) | Card(suit)
			i++
		}
	}
	return cards
}

// NewDeck52 returns card's deck from Two to Ace by all suits
func NewDeck52() Cards {
	cards, i := make(Cards, 52), 0
	for rank := Two; rank <= Ace; rank++ {
		for suit := Hearts; suit <= Clubs; suit += 1 << 4 {
			cards[i] = Card(rank) | Card(suit)
			i++
		}
	}
	return cards
}

// NewDeck54 returns card's deck from Two to Ace by all suits and Red and Black Jokers
func NewDeck54() Cards {
	cards, i := make(Cards, 54), 0
	for rank := Two; rank <= Ace; rank++ {
		for suit := Hearts; suit <= Clubs; suit += 1 << 4 {
			cards[i] = Card(rank) | Card(suit)
			i++
		}
	}
	cards[52], cards[53] = JokerRed, JokerBlack
	return cards
}
