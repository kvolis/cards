package main

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
	if PrintMode == Full {
		return suitName[suit]
	}
	if PrintMode == Short {
		return suitName[suit][:1]
	}
	return suitSymbols[suit]
}

// String returns a string representation
func (color Color) String() string {
	return colorName[color]
}
