package main

// Deck, Hand, Table, Discard, Cards are similar entities, they are separated for easy use
// Field's hiding provides safety
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
	Cards struct {
		cards cards
	}
)

type cards []Card

func (toList cards) list() Cards {
	res := make(cards, len(toList))
	for i, card := range toList {
		res[i] = card
	}
	return Cards{cards: res}
}

// NewDeck36 returns card's deck from Six to Ace by all suits
func NewDeck36() Deck {
	return Deck{cards: newDeck36()}
}

// NewDeck52 returns card's deck from Two to Ace by all suits
func NewDeck52() Deck {
	return Deck{cards: newDeck52()}
}

// NewHand returns empty entity as Hand type
func NewHand() Hand {
	return Hand{cards: emptyDeck()}
}

// NewTable returns empty entity as Table type
func NewTable() Table {
	return Table{cards: emptyDeck()}
}

// NewDiscard returns empty entity as Discard type
func NewDiscard() Discard {
	return Discard{cards: emptyDeck()}
}

// List returns deck's cards
func (deck *Deck) List() Cards {
	return deck.cards.list()
}

// List returns hands's cards
func (hand *Hand) List() Cards {
	return hand.cards.list()
}

// List returns tables's cards
func (table *Table) List() Cards {
	return table.cards.list()
}

// List returns dicard's cards
func (discard *Discard) List() Cards {
	return discard.cards.list()
}

// List returns dicard's cards
func (cards *Cards) List() Cards {
	return cards.cards.list()
}

// Give returns Deck's cards (if any) and remove it from Deck
func (deck *Deck) Give(toGive Cards) Cards {
	return Cards{cards: deck.give(toGive.cards)}
}

// Give returns Hand's cards (if any) and remove it from Hand
func (hand *Hand) Give(toGive Cards) Cards {
	return Cards{cards: hand.give(toGive.cards)}
}

// Give returns Table's cards (if any) and remove it from Table
func (table *Table) Give(toGive Cards) Cards {
	return Cards{cards: table.give(toGive.cards)}
}

// Give returns Discard's cards (if any) and remove it from Discard
func (discard *Discard) Give(toGive Cards) Cards {
	return Cards{cards: discard.give(toGive.cards)}
}

// Give returns Cards's cards (if any) and remove it from Cards
func (cards *Cards) Give(toGive Cards) Cards {
	return Cards{cards: cards.give(toGive.cards)}
}

// Take append cards to Deck
func (deck *Deck) Take(toTake Cards) {
	deck.take(toTake.cards)
}

// Take append cards to Hand
func (hand *Hand) Take(toTake Cards) {
	hand.take(toTake.cards)
}

// Take append cards to Table
func (table *Table) Take(toTake Cards) {
	table.take(toTake.cards)
}

// Take append cards to Discard
func (discard *Discard) Take(toTake Cards) {
	discard.take(toTake.cards)
}

// Take append cards to Cards
func (cards *Cards) Take(toTake Cards) {
	cards.take(toTake.cards)
}

func newDeck36() cards {
	cards, i := make([]Card, 36), 0
	for rank := Six; rank <= Ace; rank++ {
		for suit := Hearts; suit <= Clubs; suit += 1 << 4 {
			cards[i] = newCard(Name(rank) | Name(suit))
			i++
		}
	}
	return cards
}

func newDeck52() cards {
	cards, i := make([]Card, 52), 0
	for rank := Two; rank <= Ace; rank++ {
		for suit := Hearts; suit <= Clubs; suit += 1 << 4 {
			cards[i] = newCard(Name(rank) | Name(suit))
			i++
		}
	}
	return cards
}

func emptyDeck() cards {
	return []Card{}
}
