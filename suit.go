package cards

type (
	Suit  uint8
	Color uint8
)

const (
	Hearts Suit = iota << 4
	Diamonds
	Spades
	Clubs
)

const (
	Red Color = iota << 5
	Black
)

const (
	suitMask  Suit  = 0b00110000
	colorMask Color = 0b00100000
)

var suitName map[Suit]string = map[Suit]string{
	Hearts:   "Hearts",
	Diamonds: "Diamonds",
	Spades:   "Spades",
	Clubs:    "Clubs",
}

var colorName map[Color]string = map[Color]string{
	Red:   "Red",
	Black: "Black",
}

func (suit Suit) String() string {
	return suitName[suit]
}

func (color Color) String() string {
	return colorName[color]
}

func (card *Card) Suit() Suit {
	return Suit(card.name) & suitMask
}

func (card *Card) Color() Color {
	return Color(card.name) & colorMask
}

func CardSuit(card Card) Suit {
	return card.Suit()
}

func CardColor(card Card) Color {
	return card.Color()
}

func (card *Card) IsHearts() bool {
	return card.Suit() == Hearts
}

func (card *Card) IsDiamonds() bool {
	return card.Suit() == Diamonds
}

func (card *Card) IsSpades() bool {
	return card.Suit() == Spades
}

func (card *Card) IsClubs() bool {
	return card.Suit() == Clubs
}
