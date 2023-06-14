package cards

var PrintMode = Full

type printMode uint8

const (
	Full printMode = iota
	Short
	Symbol
)

const (
	TwoHearts   Card = Card(Two) | Card(Hearts)
	ThreeHearts Card = Card(Three) | Card(Hearts)
	FourHearts  Card = Card(Four) | Card(Hearts)
	FiveHearts  Card = Card(Five) | Card(Hearts)
	SixHearts   Card = Card(Six) | Card(Hearts)
	SevenHearts Card = Card(Seven) | Card(Hearts)
	EightHearts Card = Card(Eight) | Card(Hearts)
	NineHearts  Card = Card(Nine) | Card(Hearts)
	TenHearts   Card = Card(Ten) | Card(Hearts)
	JackHearts  Card = Card(Jack) | Card(Hearts)
	QueenHearts Card = Card(Queen) | Card(Hearts)
	KingHearts  Card = Card(King) | Card(Hearts)
	AceHearts   Card = Card(Ace) | Card(Hearts)

	TwoDiamonds   Card = Card(Two) | Card(Diamonds)
	ThreeDiamonds Card = Card(Three) | Card(Diamonds)
	FourDiamonds  Card = Card(Four) | Card(Diamonds)
	FiveDiamonds  Card = Card(Five) | Card(Diamonds)
	SixDiamonds   Card = Card(Six) | Card(Diamonds)
	SevenDiamonds Card = Card(Seven) | Card(Diamonds)
	EightDiamonds Card = Card(Eight) | Card(Diamonds)
	NineDiamonds  Card = Card(Nine) | Card(Diamonds)
	TenDiamonds   Card = Card(Ten) | Card(Diamonds)
	JackDiamonds  Card = Card(Jack) | Card(Diamonds)
	QueenDiamonds Card = Card(Queen) | Card(Diamonds)
	KingDiamonds  Card = Card(King) | Card(Diamonds)
	AceDiamonds   Card = Card(Ace) | Card(Diamonds)

	TwoSpades   Card = Card(Two) | Card(Spades)
	ThreeSpades Card = Card(Three) | Card(Spades)
	FourSpades  Card = Card(Four) | Card(Spades)
	FiveSpades  Card = Card(Five) | Card(Spades)
	SixSpades   Card = Card(Six) | Card(Spades)
	SevenSpades Card = Card(Seven) | Card(Spades)
	EightSpades Card = Card(Eight) | Card(Spades)
	NineSpades  Card = Card(Nine) | Card(Spades)
	TenSpades   Card = Card(Ten) | Card(Spades)
	JackSpades  Card = Card(Jack) | Card(Spades)
	QueenSpades Card = Card(Queen) | Card(Spades)
	KingSpades  Card = Card(King) | Card(Spades)
	AceSpades   Card = Card(Ace) | Card(Spades)

	TwoClubs   Card = Card(Two) | Card(Clubs)
	ThreeClubs Card = Card(Three) | Card(Clubs)
	FourClubs  Card = Card(Four) | Card(Clubs)
	FiveClubs  Card = Card(Five) | Card(Clubs)
	SixClubs   Card = Card(Six) | Card(Clubs)
	SevenClubs Card = Card(Seven) | Card(Clubs)
	EightClubs Card = Card(Eight) | Card(Clubs)
	NineClubs  Card = Card(Nine) | Card(Clubs)
	TenClubs   Card = Card(Ten) | Card(Clubs)
	JackClubs  Card = Card(Jack) | Card(Clubs)
	QueenClubs Card = Card(Queen) | Card(Clubs)
	KingClubs  Card = Card(King) | Card(Clubs)
	AceClubs   Card = Card(Ace) | Card(Clubs)
)
