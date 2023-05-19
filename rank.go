package cards

type Rank uint8

// 4 least bits define rank
const (
	Two Rank = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
	Joker
)

const rankMask = 0b00001111

var rankName map[Rank]string = map[Rank]string{
	Two:   "Two",
	Three: "Three",
	Four:  "Four",
	Five:  "Five",
	Six:   "Six",
	Seven: "Seven",
	Eight: "Eight",
	Nine:  "Nine",
	Ten:   "Ten",
	Jack:  "Jack",
	Queen: "Queen",
	King:  "King",
	Ace:   "Ace",
	Joker: "Joker",
}

func (rank Rank) String() string {
	return rankName[rank]
}

func (card *Card) Rank() Rank {
	return Rank(card.name) & rankMask
}

func CardRank(card Card) Rank {
	return card.Rank()
}

func (card *Card) IsTwo() bool {
	return card.Rank() == Two
}

func (card *Card) IsThree() bool {
	return card.Rank() == Three
}

func (card *Card) IsFour() bool {
	return card.Rank() == Four
}

func (card *Card) IsFive() bool {
	return card.Rank() == Five
}

func (card *Card) IsSix() bool {
	return card.Rank() == Six
}

func (card *Card) IsSeven() bool {
	return card.Rank() == Seven
}

func (card *Card) IsEight() bool {
	return card.Rank() == Eight
}

func (card *Card) IsNine() bool {
	return card.Rank() == Nine
}

func (card *Card) IsTen() bool {
	return card.Rank() == Ten
}

func (card *Card) IsJack() bool {
	return card.Rank() == Jack
}

func (card *Card) IsQueen() bool {
	return card.Rank() == Queen
}

func (card *Card) IsKing() bool {
	return card.Rank() == King
}

func (card *Card) IsAce() bool {
	return card.Rank() == Ace
}

func (card *Card) IsJoker() bool {
	return card.Rank() == Joker
}

func CardIsTwo(card Card) bool {
	return card.Rank() == Two
}

func CardIsThree(card Card) bool {
	return card.Rank() == Three
}

func CardIsFour(card Card) bool {
	return card.Rank() == Four
}

func CardIsFive(card Card) bool {
	return card.Rank() == Five
}

func CardIsSix(card Card) bool {
	return card.Rank() == Six
}

func CardIsSeven(card Card) bool {
	return card.Rank() == Seven
}

func CardIsEight(card Card) bool {
	return card.Rank() == Eight
}

func CardIsNine(card Card) bool {
	return card.Rank() == Nine
}

func CardIsTen(card Card) bool {
	return card.Rank() == Ten
}

func CardIsJack(card Card) bool {
	return card.Rank() == Jack
}

func CardIsQueen(card Card) bool {
	return card.Rank() == Queen
}

func CardIsKing(card Card) bool {
	return card.Rank() == King
}

func CardIsAce(card Card) bool {
	return card.Rank() == Ace
}

func CardIsJoker(card Card) bool {
	return card.Rank() == Joker
}
