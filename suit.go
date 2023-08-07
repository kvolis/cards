package cards

type (
	Suit  uint8
	Color uint8
)

// 4th and 5th bits define suit
// 5th bit enough to color define
const (
	suitMask  Suit  = 0b00110000
	colorMask Color = 0b00100000
)

// Red and Black as JRed and JBlack are used as a suit only for Jokers
// Suit(JRed) and Suit(JBlack) are not the same as Color(Red) and Color(Black)
const (
	Hearts Suit = iota << 4
	Diamonds
	Spades
	Clubs

	JRed   = Suit(Red)
	JBlack = Suit(Black)
)

// Color(Red) and Color(Black) are not the same as Suit(JRed) and Suit(JBlack)
const (
	Red Color = iota << 5
	Black
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

var suitSymbols map[Suit]string = map[Suit]string{
	Hearts:   "\u2764",
	Diamonds: "\u2666",
	Spades:   "\u2660",
	Clubs:    "\u2663",
}

// String returns a string representation
func (suit Suit) String() string {
	switch PrintMode {
	case Short:
		return suitName[suit][:1]
	case Symbol:
		return suitSymbols[suit]
	default:
		return suitName[suit]
	}
}

// String returns a string representation
func (color Color) String() string {
	if PrintMode < Full {
		return colorName[color][:1]
	}
	return colorName[color]
}
