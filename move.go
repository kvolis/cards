package cards

func (deck *Deck) give(cards []Card) []Card {
	res := make([]Card, 0, len(cards))
	for _, card := range cards {
		for i, card_ := range deck.cards {
			if card == card_ {
				res = append(res, card)
				if l := len(deck.cards); i == l-1 {
					if l == 1 {
						deck.cards = []Card{}
						return res
					}
					deck.cards = deck.cards[:i]
					break
				}
				deck.cards = append(deck.cards[:i], deck.cards[i+1:]...)
				break
			}
		}
	}
	return res
}

func (deck *Deck) take(cards []Card) {
	deck.cards = append(deck.cards, cards...)
}

func (hand *Hand) give(cards []Card) []Card {
	res := make([]Card, 0, len(cards))
	for _, card := range cards {
		for i, card_ := range hand.cards {
			if card == card_ {
				res = append(res, card)
				if l := len(hand.cards); i == l-1 {
					if l == 1 {
						hand.cards = []Card{}
						return res
					}
					hand.cards = hand.cards[:i]
					break
				}
				hand.cards = append(hand.cards[:i], hand.cards[i+1:]...)
				break
			}
		}
	}
	return res
}

func (hand *Hand) take(cards []Card) {
	hand.cards = append(hand.cards, cards...)
}

func (table *Table) give(cards []Card) []Card {
	res := make([]Card, 0, len(cards))
	for _, card := range cards {
		for i, card_ := range table.cards {
			if card == card_ {
				res = append(res, card)
				if l := len(table.cards); i == l-1 {
					if l == 1 {
						table.cards = []Card{}
						return res
					}
					table.cards = table.cards[:i]
					break
				}
				table.cards = append(table.cards[:i], table.cards[i+1:]...)
				break
			}
		}
	}
	return res
}

func (table *Table) take(cards []Card) {
	table.cards = append(table.cards, cards...)
}

func (discard *Discard) give(cards []Card) []Card {
	res := make([]Card, 0, len(cards))
	for _, card := range cards {
		for i, card_ := range discard.cards {
			if card == card_ {
				res = append(res, card)
				if l := len(discard.cards); i == l-1 {
					if l == 1 {
						discard.cards = []Card{}
						return res
					}
					discard.cards = discard.cards[:i]
					break
				}
				discard.cards = append(discard.cards[:i], discard.cards[i+1:]...)
				break
			}
		}
	}
	return res
}

func (discard *Discard) take(cards []Card) {
	discard.cards = append(discard.cards, cards...)
}
