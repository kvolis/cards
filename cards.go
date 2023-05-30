package main

import "fmt"

// Card is the simplest entity
// Field's hiding provides safety
type Card struct {
	name Name
}

// Name returns card's name as Name type
func (card Card) Name() Name {
	return card.name
}

// Rank returns card's rank as Rank type
func (card Card) Rank() Rank {
	return card.name.rank()
}

// Suit returns card's suit as Suit type
func (card Card) Suit() Suit {
	return card.name.suit()
}

// Color returns card's color as Color type
func (card Card) Color() Color {
	return card.name.color()
}

func (card Card) String() string {
	return card.Rank().String() + " " + card.Suit().String()
}

func newCard(name Name) Card {
	return Card{name: name}
}

func main() {
	deck := NewDeck36()
	fmt.Println("deck", deck)
	fmt.Println()

	hand := NewHand()
	table := NewTable()

	toMove := deck.Give(deck.List()[20:])
	fmt.Println("deck", deck)
	fmt.Println()

	hand.Take(toMove)
	fmt.Println("hand", hand)
	fmt.Println()

	toMove = hand.Give(hand.List()[10:])
	fmt.Println("hand", hand)
	fmt.Println()

	table.Take(toMove)
	fmt.Println("table", table)
	fmt.Println()

	toMove = deck.Give(deck.List()[10:])
	fmt.Println("deck", deck)
	fmt.Println()

	table.Take(toMove)
	fmt.Println("table", table)
	fmt.Println()
}
