package cards

import (
	"math/rand"
	"time"
)

// Cards are collect of Card entity
type Cards []Card

func (c Cards) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(c), func(i, j int) { c[i], c[j] = c[j], c[i] })
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
