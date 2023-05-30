package main

func (cards *cards) give(toGive cards) ([]Card, []Card) {
	res := make([]Card, 0, len(toGive))
	for _, card := range toGive {
		for i, card_ := range *cards {
			if card == card_ {
				res = append(res, card)
				if l := len(*cards); i == l-1 {
					if l == 1 {
						*cards = []Card{}
						return res, *cards
					}
					*cards = (*cards)[:i]
					break
				}
				*cards = append((*cards)[:i], (*cards)[i+1:]...)
				break
			}
		}
	}
	return res, *cards
}

func (cards *cards) take(toTake cards) {
	*cards = append(*cards, toTake...)
}

func (deck *Deck) give(toGive cards) cards {
	given, newDeck := deck.cards.give(toGive)
	deck.cards = newDeck
	return given
}

func (deck *Deck) take(toTake cards) {
	deck.cards.take(toTake)
}

func (hand *Hand) give(toGive cards) cards {
	given, newHand := hand.cards.give(toGive)
	hand.cards = newHand
	return given
}

func (hand *Hand) take(toTake cards) {
	hand.cards.take(toTake)
}

func (table *Table) give(toGive cards) cards {
	given, newTable := table.cards.give(toGive)
	table.cards = newTable
	return given
}

func (table *Table) take(toTake cards) {
	table.cards.take(toTake)
}

func (discard *Discard) give(toGive cards) cards {
	given, newDiscard := discard.cards.give(toGive)
	discard.cards = newDiscard
	return given
}

func (discard *Discard) take(toTake cards) {
	discard.cards.take(toTake)
}

func (cards *Cards) give(toGive cards) cards {
	given, newCards := cards.cards.give(toGive)
	cards.cards = newCards
	return given
}

func (cards *Cards) take(toTake cards) {
	cards.cards.take(toTake)
}
