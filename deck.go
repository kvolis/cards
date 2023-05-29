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
func (deck *Deck) List() []Card {
	return deck.cards.list()
}

// List returns hands's cards
func (hand *Hand) List() []Card {
	return hand.cards.list()
}

// List returns tables's cards
func (table *Table) List() []Card {
	return table.cards.list()
}

// List returns dicard's cards
func (discard *Discard) List() []Card {
	return discard.cards.list()
}

// Give returns deck's cards (if any) and remove it from deck
func (deck *Deck) Give(cards []Card) []Card {
	return deck.give(cards)
}

// Give returns hand's cards (if any) and remove it from hand
func (hand *Hand) Give(cards []Card) []Card {
	return hand.give(cards)
}

// Give returns table's cards (if any) and remove it from table
func (table *Table) Give(cards []Card) []Card {
	return table.give(cards)
}

// Give returns discard's cards (if any) and remove it from discard
func (discard *Discard) Give(cards []Card) []Card {
	return discard.give(cards)
}

// Take append cards to deck
func (deck *Deck) Take(cards []Card) {
	deck.take(cards)
}

// Take append cards to hand
func (hand *Hand) Take(cards []Card) {
	hand.take(cards)
}

// Take append cards to table
func (table *Table) Take(cards []Card) {
	table.take(cards)
}

// Take append cards to discard
func (discard *Discard) Take(cards []Card) {
	discard.take(cards)
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
