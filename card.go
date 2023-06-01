package cards

// Card is the simplest entity
// Field's hiding provides safety
type Card struct {
	name Name
}

// Name returns card's name as Name type
func (c Card) Name() Name {
	return c.name
}

// Rank returns card's rank as Rank type
func (c Card) Rank() Rank {
	return c.name.rank()
}

// Suit returns card's suit as Suit type
func (c Card) Suit() Suit {
	return c.name.suit()
}

// Color returns card's color as Color type
func (c Card) Color() Color {
	return c.name.color()
}

func (c Card) String() string {
	return c.name.String()
}

func newCard(name Name) Card {
	return Card{name: name}
}
