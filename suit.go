package cards

type (
	Suit  uint8
	Color uint8
)

// 4th and 5th bits define suit
const (
	Hearts Suit = iota << 4
	Diamonds
	Spades
	Clubs
)

// 5th bit enough to color define
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

func (card Card) Suit() Suit {
	return Suit(card.name) & suitMask
}

func (card Card) Color() Color {
	return Color(card.name) & colorMask
}

func CardSuit(card Card) Suit {
	return card.Suit()
}

func CardColor(card Card) Color {
	return card.Color()
}

func (card Card) IsHearts() bool {
	return card.Suit() == Hearts
}

func (card Card) IsDiamonds() bool {
	return card.Suit() == Diamonds
}

func (card Card) IsSpades() bool {
	return card.Suit() == Spades
}

func (card Card) IsClubs() bool {
	return card.Suit() == Clubs
}

func (card Card) IsRed() bool {
	return card.Color() == Red
}

func (card Card) IsBlack() bool {
	return card.Color() == Black
}

func CardIsHearts(card Card) bool {
	return card.Suit() == Hearts
}

func CardIsDiamonds(card Card) bool {
	return card.Suit() == Diamonds
}

func CardIsSpades(card Card) bool {
	return card.Suit() == Spades
}

func CardIsClubs(card Card) bool {
	return card.Suit() == Clubs
}

func CardIsRed(card Card) bool {
	return card.Color() == Red
}

func CardIsBlack(card Card) bool {
	return card.Color() == Black
}
