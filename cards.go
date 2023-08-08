package cards

import (
	"math/rand"
	"sort"
	"strings"
	"time"
)

// Cards are collect of Card entity
type Cards []Card

// Len returns cards count in the collect
func (c Cards) Len() int {
	return len(c)
}

// Shuffle shuffles cards randomly
func (c Cards) Shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(c.Len(), func(i, j int) { c[i], c[j] = c[j], c[i] })
}

// Sort sorts cards by rank from low to high and by suit (Hearts, Diamonds, Spades, Clubs)
func (c Cards) Sort() {
	sort.Slice(c, func(i, j int) bool {
		if c[i].Rank() == c[j].Rank() {
			return c[i].Suit() < c[j].Suit()
		}
		return c[i].Rank() < c[j].Rank()
	})
}

// Random returns count random cards from the collect.
// Sorting cards result is correct. Incorrect count processed normally
func (c Cards) Random(count int) Cards {
	if count < 1 {
		return Cards{}
	}
	l := c.Len()
	if count > l {
		count = l
	}

	srcIdxs := make([]int, l)
	rndIdxs := make([]int, count)
	for i := 0; i < l; i++ {
		srcIdxs[i] = i
	}
	for i := 0; i < count; i++ {
		l := len(srcIdxs)
		r := rand.Intn(l)
		rndIdxs[i] = srcIdxs[r]
		srcIdxs[r] = srcIdxs[l-1]
		srcIdxs = srcIdxs[:l-1]
	}
	sort.Ints(rndIdxs)

	res := make(Cards, count)
	for n, i := range rndIdxs {
		res[n] = c[i]
	}

	return res
}

// AreRed indicates that all cards in the collect are red
func (c Cards) AreRed() bool {
	for _, card := range c {
		if card.IsBlack() {
			return false
		}
	}
	return true
}

// AreBlack indicates that all cards in the collect are black
func (c Cards) AreBlack() bool {
	for _, card := range c {
		if card.IsRed() {
			return false
		}
	}
	return true
}

// Red returns all red cards
func (c Cards) Red() Cards {
	res := Cards{}
	for _, card := range c {
		if card.Color() == Red {
			res = append(res, card)
		}
	}
	return res
}

// Vlack returns all black cards
func (c Cards) Black() Cards {
	res := Cards{}
	for _, card := range c {
		if card.Color() == Black {
			res = append(res, card)
		}
	}
	return res
}

// AreSameSuit indicates that all cards in the collect are same suit, and returns that suit, else returns false and nil
func (c Cards) AreSameSuit() (bool, *Suit) {
	suit := c[0].Suit()
	if c.Len() == 1 {
		return true, &suit
	}

	for i := 1; i < c.Len(); i++ {
		if c[i].Suit() != suit {
			return false, nil
		}
	}

	return true, &suit
}

// Higher returns one or more cards of the highest rank
func (c Cards) Higher() Cards {
	res := Cards{c[0]}
	if c.Len() == 1 {
		return res
	}

	for i := 1; i < c.Len(); i++ {
		if c[i].Rank() < res[0].Rank() {
			continue
		}

		if c[i].Rank() > res[0].Rank() {
			res = Cards{c[i]}
			continue
		}

		res = append(res, c[i])
	}

	return res
}

// Lower returns one or more cards of the lowest rank
func (c Cards) Lower() Cards {
	res := Cards{c[0]}
	if c.Len() == 1 {
		return res
	}

	for i := 1; i < c.Len(); i++ {
		if c[i].Rank() > res[0].Rank() {
			continue
		}

		if c[i].Rank() < res[0].Rank() {
			res = Cards{c[i]}
			continue
		}

		res = append(res, c[i])
	}

	return res
}

// Shift draws top count cards and moves down.
// Incorrect count processed normally, count < 0 ignored, count > len(Cards) processed like repeat
func (c Cards) Shift(count int) {
	l := c.Len()
	if count = count % l; count == 0 || count < 0 {
		return
	}

	temp := make(Cards, count)
	copy(temp, c[l-count:])
	copy(c, c[:l-count])
	for i := 0; i < count; i++ {
		c[i] = temp[i]
	}
}

// String returns a string representation
func (c Cards) String() string {
	res := make([]string, c.Len())
	for i, card := range c {
		res[i] = card.String()
	}
	return strings.Join(res, ", ")
}

// HigherBy returns higher card by suit if any, and index, else returns nil and -1
func (c Cards) HigherBy(suit Suit) (*Card, int) {
	var (
		resCard  Card
		resIndex int = -1
	)

	for i, card := range c {
		if card.Suit() != suit {
			continue
		}
		if resIndex == -1 || card.Rank() > resCard.Rank() {
			resCard, resIndex = card, i
		}
	}

	if resIndex == -1 {
		return nil, resIndex
	}
	return &resCard, resIndex
}

// LowerBy returns higher card by suit if any, and index, else returns nil and -1
func (c Cards) LowerBy(suit Suit) (*Card, int) {
	var (
		resCard  Card
		resIndex int = -1
	)

	for i, card := range c {
		if card.Suit() != suit {
			continue
		}
		if resIndex == -1 || card.Rank() < resCard.Rank() {
			resCard, resIndex = card, i
		}
	}

	if resIndex == -1 {
		return nil, resIndex
	}
	return &resCard, resIndex
}

// By returns all cards by rank or suit or color
func (c Cards) By(rankSuitColor interface{}) Cards {
	switch rankSuitColor.(type) {
	case Rank:
		return c.byRank(rankSuitColor.(Rank))
	case Suit:
		return c.bySuit(rankSuitColor.(Suit))
	case Color:
		return c.byColor(rankSuitColor.(Color))
	default:
		return Cards{}
	}
}

// IndexBy returns index of card in cards if any, else returns -1
func (c Cards) IndexBy(card Card) int {
	for i, card_ := range c {
		if card_ == card {
			return i
		}
	}
	return -1
}

// TransferTo moves cards (if any) from src to dst.
// If there are duplicate cards in the src and transferred, they will be processed correctly, the order in the src and dst will be preserved
func (src *Cards) TransferTo(dst *Cards, transferred Cards) {
	toTransfer := make(map[Card]int)
	for _, card := range transferred {
		toTransfer[card]++
	}

	newSrc := Cards{}
	temp := make(Cards, dst.Len())
	copy(temp, *dst)
	for _, card := range *src {
		if v, ok := toTransfer[card]; ok {
			temp = append(temp, card)
			if v == 1 {
				delete(toTransfer, card)
				continue
			}
			toTransfer[card]--
			continue
		}
		newSrc = append(newSrc, card)
	}

	*src = make(Cards, len(newSrc))
	copy(*src, newSrc)

	*dst = make(Cards, len(temp))
	copy(*dst, temp)
}

// Top returns count top cards
func (c Cards) Top(count int) Cards {
	if count < 1 {
		return Cards{}
	}
	l := c.Len()
	if count > l {
		count = l
	}
	return c[:count]
}

// Bottom returns count bottom cards
func (c Cards) Bottom(count int) Cards {
	if count < 1 {
		return Cards{}
	}
	l := c.Len()
	if count > l {
		count = l
	}
	return c[l-count:]
}

// AreUnique indicates that all cards in the collect are unique
func (c Cards) AreUnique() bool {
	return c.Len() == c.Unique().Len()
}

// Unique returns unique cards (removes dublicates is any).
// The sorting of the original order is preserved
func (c Cards) Unique() Cards {
	temp := make(map[Card]struct{})
	res := Cards{}
	for _, card := range c {
		if _, ok := temp[card]; ok {
			continue
		}
		temp[card] = struct{}{}
		res = append(res, card)
	}
	return res[:]
}

func (c Cards) byRank(rank Rank) Cards {
	res := Cards{}
	for _, card := range c {
		if card.Rank() == rank {
			res = append(res, card)
		}
	}
	return res
}

func (c Cards) bySuit(suit Suit) Cards {
	res := Cards{}
	for _, card := range c {
		if card.Suit() == suit {
			res = append(res, card)
		}
	}
	return res
}

func (c Cards) byColor(color Color) Cards {
	res := Cards{}
	for _, card := range c {
		if card.Color() == color {
			res = append(res, card)
		}
	}
	return res
}
