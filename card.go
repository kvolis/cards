package cards

type Card uint8

// Rank returns card's rank
func (c Card) Rank() Rank {
	return Rank(c) & rankMask
}

// Suit returns card's suit
func (c Card) Suit() Suit {
	if c.Rank() == Joker {
		return Suit(Color(c) & colorMask)
	}
	return Suit(c) & suitMask
}

// Color returns card's color
func (c Card) Color() Color {
	return Color(c) & colorMask
}

// String returns a string representation
func (c Card) String() string {
	sep := ""
	if PrintMode == Full {
		sep = " "
	}

	suit := c.Suit().String()
	rank := c.Rank()
	if rank == Joker {
		suit = c.Color().String()
	}

	return rank.String() + sep + suit
}

// IsRed indicates that card is red
func (c Card) IsRed() bool {
	return c.Color() == Red
}

// IsBlack indicates that card is black
func (c Card) IsBlack() bool {
	return c.Color() == Black
}