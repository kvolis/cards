package cards

import "strings"

// Cards are collect of Card entity
// Field's hiding provides safety
type Cards struct {
	cards []Card
}

// List returns list of Card entities
func (c *Cards) List() []Card {
	res := make([]Card, len(c.cards))
	for i, card := range c.cards {
		res[i] = card
	}
	return res
}

// Give returns Cards by indexes and remove it from Cards
// Incorrect (i < 0 | i >= len(Cards)) indexes are ignored
// Non-unique index be processed once
func (c *Cards) Give(indices ...int) Cards {
	uniq := map[int]struct{}{}
	l := len(c.cards)
	for _, i := range indices {
		if i >= 0 && i < l {
			uniq[i] = struct{}{}
		}
	}

	res, remain := []Card{}, []Card{}
	for i, card := range c.cards {
		if _, ok := uniq[i]; !ok {
			remain = append(remain, card)
			continue
		}
		res = append(res, card)
	}

	c.cards = remain
	return Cards{cards: res}
}

// GiveAll returns Cards from source and flush source
func (c *Cards) GiveAll() Cards {
	var res Cards
	res.cards, c.cards = c.cards, res.cards
	return res
}

// GiveTop returns count top cards and remove it from Cards
// Incorrect count processed normally, it returns empty if count<=0 or full if count>len(Cards)
func (c *Cards) GiveTop(count int) Cards {
	if count <= 0 {
		return EmptyCards()
	}
	if count > len(c.cards) {
		return c.GiveAll()
	}
	var res Cards
	res.cards, c.cards = c.cards[:count], c.cards[count:]
	return res
}

// GiveBottom returns count bottom cards and remove it from Cards
// Incorrect count processed normally, it returns empty if count<=0 or full if count>len(Cards)
func (c *Cards) GiveBottom(count int) Cards {
	if count <= 0 {
		return EmptyCards()
	}
	var l int
	if l = len(c.cards); count > l {
		return c.GiveAll()
	}
	var res Cards
	res.cards, c.cards = c.cards[l-count:], c.cards[:l-count]
	return res
}

// Take appends cards to Cards and flush gaven
func (c *Cards) Take(cards *Cards) {
	c.cards = append(c.cards, cards.cards...)
	*cards = EmptyCards()
}

func (c *Cards) Len() int {
	return len(c.cards)
}

func (c Cards) String() string {
	res := make([]string, len(c.cards))
	for i, card := range c.cards {
		res[i] = card.name.String()
	}
	return strings.Join(res, ", ")
}
