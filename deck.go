package cards

// Deck, Hand, Table, Discard are similar entities, they are separated for easy use
type (
	Deck struct {
		cards cards
	}
	Hand struct {
		cards cards
	}
	Table struct {
		cards cards
	}
	Discard struct {
		cards cards
	}
)

type cards []Card

func (cards cards) list() []Card {
	res := make([]Card, len(cards))
	for i, card := range cards {
		res[i] = card
	}
	return res
}

func NewDeck36() Deck {
	return Deck{cards: newDeck36()}
}

func NewDeck52() Deck {
	return Deck{cards: newDeck52()}
}

func NewHand() Hand {
	return Hand{cards: emptyDeck()}
}

func NewTable() Table {
	return Table{cards: emptyDeck()}
}

func NewDiscard() Discard {
	return Discard{cards: emptyDeck()}
}

func (deck *Deck) List() []Card {
	return deck.cards.list()
}

func (hand *Hand) List() []Card {
	return hand.cards.list()
}

func (table *Table) List() []Card {
	return table.cards.list()
}

func (discard *Discard) List() []Card {
	return discard.cards.list()
}

func (deck *Deck) Give(cards []Card) []Card {
	return deck.give(cards)
}

func (hand *Hand) Give(cards []Card) []Card {
	return hand.give(cards)
}

func (table *Table) Give(cards []Card) []Card {
	return table.give(cards)
}

func (discard *Discard) Give(cards []Card) []Card {
	return discard.give(cards)
}

func (deck *Deck) Take(cards []Card) {
	deck.take(cards)
}

func (hand *Hand) Take(cards []Card) {
	hand.take(cards)
}

func (table *Table) Take(cards []Card) {
	table.take(cards)
}

func (discard *Discard) Take(cards []Card) {
	discard.take(cards)
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

func emptyDeck() []Card {
	return []Card{}
}
