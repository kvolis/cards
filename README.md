# Cards
##### Simple and easy to create card games

This package is handy for creating most of the popular card games. Its functionality is sufficient for quick and easy operations with entities such as Card and Cards (such as Deck, Hand, Discard, etc).

Here are the following entities that you can operate on:
- `Rank` (from **Two** to **Ace**)
- `Suit` (**Hearts**, **Diamonds**, **Spades**, **Crosses**) (and additionally **JRed** as Red and **JBlack** as Black for Jokers, because their suit is determined only by color)
- `Color` (**Red**, **Black**)
- `Card` (all combinations of `Rank` and `Suit`)
- `Cards` (collection of `Card` as a standard Go slice)

All entities satisfy the standard library **Stringer** interface, so you can pretty print its. You can change the print mode through the variable `PrintMode` with the valid values **Short**, **Symbol**, **Full**.

## Features
- Create decks of 36, 52 or 54 cards (Jokers available)
- Shuffle cards randomly
- Sorting cards
- Getting cards from any collect by `Rank`, `Suit`, `Color`
- Getting lower and higher cards from the collect
- and others, see the full description of the methods below

###### Note that
the `Card` entity is a `uint8` type. The _lower 4 bits_ determine the `Rank` of the `Card`, the next _2 bits_ determine its `Suit`, with the _5th bit_ indicating its `Color`. Thus, both `Card` and `Cards` take up little space in memory.

This gives you _2 free bits_ in which you can store additional information about a particular card (for example, the card is face up or face down on the table).

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

## Documentation
```go
deck := cards.NewDeck36() // sorted deck of 36 cards
cards.NewDeck52()         // same with 52 cards
cards.NewDeck54()         // and with Jokers
// deck = {Six Hearts, Six Diamonds, ..., Ace Spades, Ace Clubs}

deck.Shuffle()
// deck = {Jack Hearts, Seven Spades, ..., Six Hearts, Queen Hearts, Seven Clubs, Nine Clubs}

deck.Shift(2) //
// deck = {Seven Clubs, Nine Clubs, Jack Hearts, Seven Spades, ..., Six Hearts, Queen Hearts}

player1, player2 := cards.Cards{}, cards.Cards{}

deck.TransferTo(&player1, deck.Top(5))
deck.TransferTo(&player2, deck.Bottom(5))
// player1 = {King Diamonds, Eight Clubs, Ace Diamonds, Six Hearts, Queen Hearts}
// player2 = {Seven Clubs, Nine Clubs, Jack Hearts, Seven Spades, Nine Spades}

h1, l1 := player1.Higher(), player1.Lower()
h2, l2 := player2.Higher(), player2.Lower()
// h1, l1 = {Ace Diamonds}, {Six Hearts}
// h2, l2 = {Jack Hearts}, {Seven Clubs, Seven Spades}

sp, spi := deck.LowerBy(cards.Spades)
// sp, spi = Six Spades, 25

player3 := cards.Cards{}
deck.TransferTo(&player3, deck.Random(5))
// player3 = {Nine Diamonds, Ten Clubs, Jack Spades, Queen Clubs, Six Diamonds}

red, black := deck.Red(), deck.Black()
// red = {Jack Diamonds, Ten Hearts, Seven Hearts, King Hearts, Eight Diamonds, Ace Hearts, ...}
// black = {King Spades, Eight Spades, Ten Spades, Jack Clubs, King Clubs, Six Clubs, ...}

spades := deck.By(cards.Spades)
sameSuit, _ := spades.AreSameSuit()
// sameSuit = true

fourDiamonds := cards.FourDiamonds

_, _, _ = fourDiamonds.Rank(), fourDiamonds.Suit(), fourDiamonds.Color()
// Four, Diamonds, Red

_, _ = fourDiamonds.IsRed(), fourDiamonds.IsBlack()
// true, false

player3.AreUnique()
// true

player3 = append(player3, cards.SevenClubs, cards.SevenClubs)
player3.AreUnique()
// false

fmt.Printf("%08b \n", player1)
// [00011011 00110110 00011100 00000100 00001010]
player1.SetHOpt(true)     // 7th, higher of two free bits = 1; indicates, for example, that the cards are face down
player1[1].SetHOpt(false) // 0 value for one element
fmt.Printf("%08b \n", player1)
// [10011011 00110110 10011100 10000100 10001010]

fmt.Printf("%08b \n", player2)
// [00110101 00110111 00001001 00100101 00100111]
player2.SetLOpt(true)     // 6th, lower of two free bits = 1, indicates, for example, that cards were viewed by the player
player2[2].SetLOpt(false) // 0 value for one element
fmt.Printf("%08b \n", player2)
// [01110101 01110111 00001001 01100101 01100111]
```

##### _Contribution are welcome_