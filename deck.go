package main

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
