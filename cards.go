package main

import (
	"math/rand"
	"sort"
	"strings"
	"time"
)

// Cards are collect of Card entity
type Cards []Card

// Shuffle shuffles cards randomly
func (c Cards) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(c), func(i, j int) { c[i], c[j] = c[j], c[i] })
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
	if len(c) == 1 {
		return true, &suit
	}

	for i := 1; i < len(c); i++ {
		if c[i].Suit() != suit {
			return false, nil
		}
	}

	return true, &suit
}

// Higher returns one or more cards of the highest rank
func (c Cards) Higher() Cards {
	res := Cards{c[0]}
	if len(c) == 1 {
		return res
	}

	for i := 1; i < len(c); i++ {
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
	if len(c) == 1 {
		return res
	}

	for i := 1; i < len(c); i++ {
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

// Shift draws top count cards and moves down
// Incorrect count processed normally, count < 0 ignored, count > len(Cards) processed like repeat
func (c Cards) Shift(count int) {
	l := len(c)
	if count = count % l; count == 0 || count < 0 {
		return
	}

	temp := make(Cards, count)
	copy(temp, c[:count])
	copy(c, c[count:])
	for i := 0; i < count; i++ {
		c[l-count+i] = temp[i]
	}
}

// String returns a string representation
func (c Cards) String() string {
	res := make([]string, len(c))
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
