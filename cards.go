package cards

// Card is the simplest entity
// Field's hiding provides safety
type Card struct {
	name Name
}

// Name returns card's name as Name type
func (card Card) Name() Name {
	return card.name
}

// Rank returns card's rank as Rank type
func (card Card) Rank() Rank {
	return card.name.rank()
}

// Suit returns card's suit as Suit type
func (card Card) Suit() Suit {
	return card.name.suit()
}

// Color returns card's color as Color type
func (card Card) Color() Color {
	return card.name.color()
}

func (card Card) String() string {
	return card.Rank().String() + " " + card.Suit().String()
}

func newCard(name Name) Card {
	return Card{name: name}
}
