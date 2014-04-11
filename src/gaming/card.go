package gaming

import (
	"fmt"
	"math/rand"
)

type Suit interface {
	Name() string
	Symbol() rune
	Index() uint
}

type suitImpl struct {
	name   string
	symbol rune
	index  uint
}

func (this *suitImpl) Name() string {
	return this.name
}

func (this *suitImpl) Symbol() rune {
	return this.symbol
}

func (this *suitImpl) Index() uint {
	return this.index
}

func (this *suitImpl) String() string {
	return this.Name()
}

var Spade = &suitImpl{name: "spade", symbol: 's', index: 0}
var Club = &suitImpl{name: "club", symbol: 'c', index: 1}
var Heart = &suitImpl{name: "heart", symbol: 'h', index: 2}
var Diamond = &suitImpl{name: "diamond", symbol: 'd', index: 3}
var suits = []Suit{Spade, Club, Heart, Diamond}

func Suits() []Suit {
	return suits
}

func RandomSuit(r *rand.Rand) Suit {
	return Suits()[r.Intn(len(suits))]
}

type Value interface {
	Name() string
	Symbol() rune
	Index() uint
}

type valueImpl struct {
	name   string
	symbol rune
	index  uint
}

func (this *valueImpl) Name() string {
	return this.name
}

func (this *valueImpl) Symbol() rune {
	return this.symbol
}

func (this *valueImpl) Index() uint {
	return this.index
}

func (this *valueImpl) String() string {
	return fmt.Sprintf("%s", this.name)
}

var Ace = &valueImpl{name: "ace", symbol: 'A', index: 0}
var Two = &valueImpl{name: "two", symbol: '2', index: 1}
var Three = &valueImpl{name: "three", symbol: '3', index: 2}
var Four = &valueImpl{name: "four", symbol: '4', index: 3}
var Five = &valueImpl{name: "five", symbol: '5', index: 4}
var Six = &valueImpl{name: "six", symbol: '6', index: 5}
var Seven = &valueImpl{name: "seven", symbol: '7', index: 6}
var Eight = &valueImpl{name: "eight", symbol: '8', index: 7}
var Nine = &valueImpl{name: "nine", symbol: '9', index: 8}
var Ten = &valueImpl{name: "ten", symbol: 'T', index: 9}
var Jack = &valueImpl{name: "jack", symbol: 'J', index: 10}
var Queen = &valueImpl{name: "queen", symbol: 'Q', index: 11}
var King = &valueImpl{name: "king", symbol: 'K', index: 12}
var Joker = &valueImpl{name: "joker", symbol: 'O', index: 13}

func Values() []Value {
	return []Value{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}
}

func JokerValues() []Value {
	return []Value{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Joker}
}

type Card interface {
	Suit() Suit
	Value() Value
}

type cardImpl struct {
	suit  Suit
	value Value
}

func (this *cardImpl) Suit() Suit {
	return this.suit
}

func (this *cardImpl) Value() Value {
	return this.value
}
