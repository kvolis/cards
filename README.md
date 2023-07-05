# Cards
##### Simple and easy to create card games

This package is handy for creating most of the popular card games. Its functionality is sufficient for quick and easy operations with entities such as Card and Cards (such as Deck, Hand, Discard, etc). There are no Jokers in the current implementation, so make sure you don't need these cards (unless you plan to cheat :-) Work with Jokers will definitely be implemented later.

Here are the following entities that you can operate on:
- `Rank` (from **Two** to **Ace**)
- `Suit` (**Hearts**, **Diamonds**, **Spades**, **Crosses**)
- `Color` (**Red**, **Black**)
- `Card` (all combinations of `Rank` and `Suit`)
- `Cards` (collection of `Card` as a standard Go slice)

All entities satisfy the standard library **Stringer** interface, so you can pretty print its.

## Features
- Create decks of 36 or 52 cards
- Shuffle cards randomly
- Sorting cards
- Getting cards from any collect by `Rank`, `Suit`, `Color`
- Getting lower and higher cards from the collect
- and others, see the full description of the methods below

###### Note that
the `Card` entity is a `uint8` type. The _lower 4 bits_ determine the `Rank` of the `Card`, the next _2 bits_ determine its `Suit`, with the _5th bit_ indicating its `Color`. Thus, both `Card` and `Cards` take up little space in memory.

This gives you _2 free bits_ in which you can store additional information about a particular card (for example, the card is face up or face down on the table). The possibility of operations with the two most significant bits will definitely be added later.

`Cards` is a collect of `Card` like a standart Go slice. Accordingly, you can define slices, add and remove elements, move cards from one slice to another - all this through standard operations when working with slices. Depending on the particular game, think of a collect of `Card` as a Deck, a Hand, a Discard, and so on - these are all representations of the same entity.

## Installation and usage
Download package
```sh
go get github.com/kvolis/cards
```
and import as usual
```go
import github.com/kvolis/cards
```