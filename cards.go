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
