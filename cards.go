package cards

type Card struct {
	name Name
}

func newCard(name Name) Card {
	return Card{name: name}
}

func (card Card) String() string {
	return card.name.String()
}

func (card Card) HigherThan(card_ Card) bool {
	return card.name.rankByName() > card_.name.rankByName()
}

func (card Card) LowerThan(card_ Card) bool {
	return card.name.rankByName() < card_.name.rankByName()
}

func (card Card) Equal(card_ Card) bool {
	return card.name.rankByName() == card_.name.rankByName()
}
