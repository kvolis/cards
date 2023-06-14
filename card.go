package cards

type Card uint8

// Rank returns card's rank
func (c Card) Rank() Rank {
	return Rank(c) & rankMask
}

// Suit returns card's suit
func (c Card) Suit() Suit {
	return Suit(c) & suitMask
}

// Color returns card's color
func (c Card) Color() Color {
	return Color(c) & colorMask
}

func (c Card) String() string {
	if PrintMode > Full {
		return c.Rank().String() + c.Suit().String()
	}
	return c.Rank().String() + " " + c.Suit().String()
}

func (c Card) IsRed() bool {
	return c.Color() == Red
}

func (c Card) IsBlack() bool {
	return c.Color() == Black
}
