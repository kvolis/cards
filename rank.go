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
}

func (rank Rank) String() string {
	if CardsPrintMode == Full {
		return rankName[rank]
	}
	if rank <= Nine {
		return string(rune(rank + 2 + '0'))
	}
	if rank >= Jack {
		return rankName[rank][:1]
	}
	return "10"
}
