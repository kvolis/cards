package main

import "fmt"

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

// String returns a string representation
func (c Card) String() string {
	if PrintMode > Full {
		return c.Rank().String() + c.Suit().String()
	}
	return c.Rank().String() + " " + c.Suit().String()
}

// IsRed indicates that card is red
func (c Card) IsRed() bool {
	return c.Color() == Red
}

// IsBlack indicates that card is black
func (c Card) IsBlack() bool {
	return c.Color() == Black
}

func main() {
	c := NewDeck36()
	c.Shuffle()
	c = c[:17]
	fmt.Println(c)
	fmt.Println()

	fmt.Println(c.IndexBy(SevenSpades))
	fmt.Println()
	c.Sort()

	fmt.Println(c)
	fmt.Println()
	fmt.Println(c.HigherBy(Hearts))
	fmt.Println(c.HigherBy(Spades))
	fmt.Println(c.HigherBy(Diamonds))
	fmt.Println(c.HigherBy(Clubs))

	fmt.Println()

	fmt.Println(c.LowerBy(Hearts))
	fmt.Println(c.LowerBy(Spades))
	fmt.Println(c.LowerBy(Diamonds))
	fmt.Println(c.LowerBy(Clubs))

	fmt.Println()
	fmt.Println(c.Higher())
	fmt.Println(c.Lower())

	fmt.Println()
	fmt.Println(c.Black())
	fmt.Println()
	fmt.Println(c.Red())

	fmt.Println()
	fmt.Println(c.By(Hearts))

	fmt.Println()
	fmt.Println(c.By(Spades))

	fmt.Println()
	fmt.Println(c.By(Ten))

	fmt.Println()
	fmt.Println(c.By(Ace))

	fmt.Println()
	fmt.Println(c.By(Queen))

	fmt.Println()
	fmt.Println(c.By(AceClubs))

	fmt.Println()
	fmt.Println(c.By(Black).Higher())

}
