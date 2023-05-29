package cards

type Card struct {
	name Name
}

func (card Card) Name() Name {
	return card.name
}

func (card Card) Rank() Rank {
	return card.name.rank()
}

func (card Card) Suit() Suit {
	return card.name.suit()
}

func (card Card) Color() Color {
	return card.name.color()
}

func newCard(name Name) Card {
	return Card{name: name}
}
