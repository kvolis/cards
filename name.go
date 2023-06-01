package cards

// Name is single property of Card
// 6 lower bits contain the rank (4) and suit (2)
// 2 higher bits can be used in the future as additional properties
type Name uint8

const (
	TwoHearts   Name = Name(Two) | Name(Hearts)
	ThreeHearts Name = Name(Three) | Name(Hearts)
	FourHearts  Name = Name(Four) | Name(Hearts)
	FiveHearts  Name = Name(Five) | Name(Hearts)
	SixHearts   Name = Name(Six) | Name(Hearts)
	SevenHearts Name = Name(Seven) | Name(Hearts)
	EightHearts Name = Name(Eight) | Name(Hearts)
	NineHearts  Name = Name(Nine) | Name(Hearts)
	TenHearts   Name = Name(Ten) | Name(Hearts)
	JackHearts  Name = Name(Jack) | Name(Hearts)
	QueenHearts Name = Name(Queen) | Name(Hearts)
	KingHearts  Name = Name(King) | Name(Hearts)
	AceHearts   Name = Name(Ace) | Name(Hearts)

	TwoDiamonds   Name = Name(Two) | Name(Diamonds)
	ThreeDiamonds Name = Name(Three) | Name(Diamonds)
	FourDiamonds  Name = Name(Four) | Name(Diamonds)
	FiveDiamonds  Name = Name(Five) | Name(Diamonds)
	SixDiamonds   Name = Name(Six) | Name(Diamonds)
	SevenDiamonds Name = Name(Seven) | Name(Diamonds)
	EightDiamonds Name = Name(Eight) | Name(Diamonds)
	NineDiamonds  Name = Name(Nine) | Name(Diamonds)
	TenDiamonds   Name = Name(Ten) | Name(Diamonds)
	JackDiamonds  Name = Name(Jack) | Name(Diamonds)
	QueenDiamonds Name = Name(Queen) | Name(Diamonds)
	KingDiamonds  Name = Name(King) | Name(Diamonds)
	AceDiamonds   Name = Name(Ace) | Name(Diamonds)

	TwoSpades   Name = Name(Two) | Name(Spades)
	ThreeSpades Name = Name(Three) | Name(Spades)
	FourSpades  Name = Name(Four) | Name(Spades)
	FiveSpades  Name = Name(Five) | Name(Spades)
	SixSpades   Name = Name(Six) | Name(Spades)
	SevenSpades Name = Name(Seven) | Name(Spades)
	EightSpades Name = Name(Eight) | Name(Spades)
	NineSpades  Name = Name(Nine) | Name(Spades)
	TenSpades   Name = Name(Ten) | Name(Spades)
	JackSpades  Name = Name(Jack) | Name(Spades)
	QueenSpades Name = Name(Queen) | Name(Spades)
	KingSpades  Name = Name(King) | Name(Spades)
	AceSpades   Name = Name(Ace) | Name(Spades)

	TwoClubs   Name = Name(Two) | Name(Clubs)
	ThreeClubs Name = Name(Three) | Name(Clubs)
	FourClubs  Name = Name(Four) | Name(Clubs)
	FiveClubs  Name = Name(Five) | Name(Clubs)
	SixClubs   Name = Name(Six) | Name(Clubs)
	SevenClubs Name = Name(Seven) | Name(Clubs)
	EightClubs Name = Name(Eight) | Name(Clubs)
	NineClubs  Name = Name(Nine) | Name(Clubs)
	TenClubs   Name = Name(Ten) | Name(Clubs)
	JackClubs  Name = Name(Jack) | Name(Clubs)
	QueenClubs Name = Name(Queen) | Name(Clubs)
	KingClubs  Name = Name(King) | Name(Clubs)
	AceClubs   Name = Name(Ace) | Name(Clubs)
)

func (name Name) String() string {
	return name.rank().String() + " " + name.suit().String()
}

func (name Name) rank() Rank {
	return Rank(name) & rankMask
}

func (name Name) suit() Suit {
	return Suit(name) & suitMask
}

func (name Name) color() Color {
	return Color(name) & colorMask
}
