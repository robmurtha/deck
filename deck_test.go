package deck

import (
	"testing"

	"strings"

	"github.com/cheekybits/is"
)

var plainCards = "SA S2 S3 S4 S5 S6 S7 S8 S9 S10 SJ SQ SK HA H2 H3 H4 H5 H6 H7 H8 H9 H10 HJ HQ HK DA D2 D3 D4 D5 D6 D7 D8 D9 D10 DJ DQ DK CA C2 C3 C4 C5 C6 C7 C8 C9 C10 CJ CQ CK"
var jokerCards = plainCards + " J"

func TestNew(t *testing.T) {
	is := is.New(t)

	d := New()
	is.Equal(52, len(d.Cards()))
	cards := []string{}
	for _, c := range d.Cards() {
		cards = append(cards, c.String())
	}
	is.Equal(plainCards, strings.Join(cards, " "))
}

func TestShuffle(t *testing.T) {
	is := is.New(t)

	d1 := New()
	d2 := New()
	is.Equal(d1.Cards(), d2.Cards())

	d1.Shuffle()
	is.NotEqual(d1.Cards(), d2.Cards())
}

func TestDeal(t *testing.T) {
	is := is.New(t)

	d := New()
	is.Equal(52, len(d.Cards()))
	is.Equal(52, cap(d.Cards()))

	_, ok := d.Deal()
	is.True(ok)
	is.Equal(51, len(d.Cards()))

	for i := 0; i < 51; i++ {
		_, ok := d.Deal()
		is.OK(ok)
		is.Equal(len(d.Cards()), cap(d.Cards()))
	}
	_, ok = d.Deal()
	is.False(ok)
}

func TestJokerDeck(t *testing.T) {
	is := is.New(t)

	var JokerDeck = Type{
		Name:   "Joker",
		Suites: []uint8{Spade, Heart, Diamond, Club},
		Values: []uint8{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King},
		InitFunc: func(dt Type) []Card {
			cards := make([]Card, 0, 53)
			cards = append(cards, GenerateCards(dt)...)
			cards = append(cards, Card{Suite: Joker, Value: None})
			return cards
		},
	}

	d := Deck{Type: JokerDeck}
	d.Reset()

	d.Shuffle()
	is.Equal(53, len(d.Cards()))
	is.Equal(53, cap(d.Cards()))
	cards := []string{}
	for _, c := range d.Cards() {
		cards = append(cards, c.String())
	}
	is.NotEqual(jokerCards, strings.Join(cards, " "))

	d.Reset()
	is.Equal(53, len(d.Cards()))
	is.Equal(53, cap(d.Cards()))
	is.Equal("Joker", d.Name)

	cards = cards[:0]
	for _, c := range d.Cards() {
		cards = append(cards, c.String())
	}
	is.Equal(jokerCards, strings.Join(cards, " "))

}

func TestInvalidDeck(t *testing.T) {
	is := is.New(t)

	var JokerDeck = Type{
		Name:   "Invalid",
		Suites: []uint8{99},
		Values: []uint8{99},
		InitFunc: GenerateCards,
	}

	d := Deck{Type: JokerDeck}
	d.Reset()
	c,_ := d.Deal()
 	is.Equal("??",c.String())
}
