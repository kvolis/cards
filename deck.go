package cards

// NewDeck36 returns card's deck from Six to Ace by all suits
func NewDeck36() Cards {
	return Cards{cards: newDeck36()}
}

// NewDeck52 returns card's deck from Two to Ace by all suits
func NewDeck52() Cards {
	return Cards{cards: newDeck52()}
}

// EmptyCards returns empty object for using as Deck, Hand, Table, etc
func EmptyCards() Cards {
	return Cards{cards: []Card{}}
}

func newDeck36() []Card {
	cards, i := make([]Card, 36), 0
	for rank := Six; rank <= Ace; rank++ {
		for suit := Hearts; suit <= Clubs; suit += 1 << 4 {
			cards[i] = newCard(Name(rank) | Name(suit))
			i++
		}
	}
	return cards
}

func newDeck52() []Card {
	cards, i := make([]Card, 52), 0
	for rank := Two; rank <= Ace; rank++ {
		for suit := Hearts; suit <= Clubs; suit += 1 << 4 {
			cards[i] = newCard(Name(rank) | Name(suit))
			i++
		}
	}
	return cards
}
