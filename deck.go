package cards

import (
	"fmt"
	"sort"
)

type Deck struct {
	deck []Card
}

func newDeck36() []Card {
	deck, i := make([]Card, 36), 0
	for rank := Six; rank <= Ace; rank++ {
		for suit := Hearts; suit <= Clubs; suit += 1 << 4 {
			deck[i] = newCard(Name(rank) | Name(suit))
			i++
		}
	}
	return deck
}

func newDeck52() []Card {
	deck, i := make([]Card, 52), 0
	for rank := Two; rank <= Ace; rank++ {
		for suit := Hearts; suit <= Clubs; suit += 1 << 4 {
			deck[i] = newCard(Name(rank) | Name(suit))
			i++
		}
	}
	return deck
}

func newDeck54() []Card {
	deck := make([]Card, 54)
	copy(deck, newDeck52())
	deck[52], deck[53] = newCard(JokerRed), newCard(JokerBlack)
	return deck
}

func NewDeck36() Deck {
	return Deck{deck: newDeck36()}
}

func NewDeck52() Deck {
	return Deck{deck: newDeck52()}
}

func NewDeck54() Deck {
	return Deck{deck: newDeck54()}
}

func (deck Deck) String() string {
	return fmt.Sprint(deck.deck)
}

func (deck Deck) Sort() {
	deck.SortByRank()
}

func (deck Deck) SortByRank() {
	sort.Sort(byRank(deck.deck))
}

func (deck Deck) SortBySuit() {
	sort.Sort(bySuit(deck.deck))
}

type (
	byRank []Card
	bySuit []Card
)

func (deck byRank) Len() int {
	return len(deck)
}

func (deck byRank) Less(i, j int) bool {
	if deck[i].name.rankByName() == deck[j].name.rankByName() {
		return deck[i].name.suitByName() < deck[j].name.suitByName()
	}
	return deck[i].name.rankByName() < deck[j].name.rankByName()
}

func (deck byRank) Swap(i, j int) {
	deck[i], deck[j] = deck[j], deck[i]
}

func (deck bySuit) Len() int {
	return len(deck)
}

func (deck bySuit) Less(i, j int) bool {
	if deck[i].name.suitByName() == deck[j].name.suitByName() {
		return deck[i].name.rankByName() < deck[j].name.rankByName()
	}
	return deck[i].name.suitByName() < deck[j].name.suitByName()
}

func (deck bySuit) Swap(i, j int) {
	deck[i], deck[j] = deck[j], deck[i]
}
