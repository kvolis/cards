package cards

type Rank uint8

// 4 least bits define rank
const rankMask = 0b00001111

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
	Joker: "Jocker",
}

// String returns a string representation
func (rank Rank) String() string {
	switch {
	case PrintMode == Full:
		return rankName[rank]
	case rank == Joker:
		return "X"
	case rank <= Nine:
		return string(rune(rank + 2 + '0'))
	case rank >= Jack:
		return rankName[rank][:1]
	default:
		return "10"
	}
}
