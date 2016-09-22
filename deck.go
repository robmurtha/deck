package deck

import (
	"math/rand"
	"time"
)

const (
	// card values
	None uint8 = iota
	Ace
	Two
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

	// card suites
	Spade
	Heart
	Diamond
	Club

	Joker
)

// suiteMap provides a string character representation of the card suite.
var suiteMap = map[uint8]string{Spade: "S", Heart: "H", Diamond: "D", Club: "C", Joker: "J"}

// valueMap provides a string character name of the card value.
var valueMap = map[uint8]string{Ace: "A", Two: "2", Three: "3", Four: "4", Five: "5", Six: "6", Seven: "7", Eight: "8", Nine: "9", Ten: "10", Jack: "J", Queen: "Q", King: "K", None: ""}

// DeckType supports the definition and creation of standard and non standard decks of cards.
type Type struct {
	Name     string
	Suites   []uint8
	Values   []uint8
	InitFunc func(Type) []Card
}

// PlainDeck represents a common deck of 52 cards.
var PlainDeck = Type{
	Name:     "Plain",
	Suites:   []uint8{Spade, Heart, Diamond, Club},
	Values:   []uint8{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King},
	InitFunc: GenerateCards,
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// Card stores the Suite and Value that uniquely identify an individual card.
// Ace through King values are mapped to 1-13.
type Card struct {
	Suite uint8
	Value uint8
}

// String provides a two character string representation of the Suite and Value fields.
func (c Card) String() string {
	s, ok := suiteMap[c.Suite]
	if !ok {
		s = "?"
	}
	v, ok := valueMap[c.Value]
	if !ok {
		v = "?"
	}
	return s + v
}

// Deck provides a concrete implementation of a mutable Card deck. Synchronization choice is left to the caller for concurrent access.
type Deck struct {
	Type
	cards []Card
}

// New is the default constructor for a plain deck of cards. Cards are not shuffled by default.
func New() *Deck {
	d := &Deck{Type: PlainDeck}
	d.Reset()
	return d
}

// Reset starts a new deck based on the DeckType.
func (d *Deck) Reset() {
	d.cards = d.InitFunc(d.Type)
}

// Cards returns an slice of cards currently in the deck.
func (d *Deck) Cards() []Card {
	return d.cards
}

// Deal returns a card and boolean, if the deck is empty the boolean will be false.
func (d *Deck) Deal() (Card, bool) {
	c := Card{}
	if len(d.cards) == 0 {
		return c, false
	}
	c, d.cards = d.cards[0], d.cards[1:]
	return c, true
}

// Shuffle randomizes the deck.
func (d *Deck) Shuffle() {
	l := len(d.cards)
	for i := range d.cards {
		r := rand.Intn(l)
		d.cards[i], d.cards[r] = d.cards[r], d.cards[i]
	}
}

// GenerateCards provides a common function for generating simple decks by combining all suites and values.
func GenerateCards(dt Type) []Card {
	slen := uint8(len(dt.Suites))
	vlen := uint8(len(dt.Values))
	cards := make([]Card, slen*vlen)

	c := 0
	for s := uint8(0); s < slen; s++ {
		for v := uint8(0); v < vlen; v++ {
			cards[c] = Card{Suite: dt.Suites[s], Value: dt.Values[v]}
			c++
		}
	}
	return cards
}
