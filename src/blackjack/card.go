package blackjack

import (
	"gaming"
	"math/rand"
)

type Value interface {
gaming.Value
	Score() uint
}

type valueImpl struct {
	gaming.Value
	score uint
}

func (this *valueImpl) Score() uint {
	return this.score
}

var Ace = &valueImpl{gaming.Ace, uint(11)}
var Two = &valueImpl{gaming.Two, uint(2)}
var Three = &valueImpl{gaming.Three, uint(3)}
var Four = &valueImpl{gaming.Four, uint(4)}
var Five = &valueImpl{gaming.Five, uint(5)}
var Six = &valueImpl{gaming.Six, uint(6)}
var Seven = &valueImpl{gaming.Seven, uint(7)}
var Eight = &valueImpl{gaming.Eight, uint(8)}
var Nine = &valueImpl{gaming.Nine, uint(9)}
var Ten = &valueImpl{gaming.Ten, uint(10)}
var Jack = &valueImpl{gaming.Jack, uint(10)}
var Queen = &valueImpl{gaming.Queen, uint(10)}
var King = &valueImpl{gaming.King, uint(10)}

var values = []Value{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

func Values() []Value {
	return values;
}

func RandomValue(r *rand.Rand) Value {
	return values[r.Intn(len(values))]
}

type Card interface {
gaming.Card
	Score() uint
}

type cardImpl struct {
	suit    gaming.Suit
	value   Value
}

func NewCard(suit gaming.Suit, value Value) Card {
	return &cardImpl{suit, value}
}

func NewRandomCard(r *rand.Rand) Card {
	return &cardImpl{suit: gaming.RandomSuit(r), value: RandomValue(r)}
}

func (this *cardImpl) Score() uint {
	return this.value.Score()
}

func (this *cardImpl) Suit() gaming.Suit {
	return this.suit
}

func (this *cardImpl) Value() gaming.Value {
	return this.value
}
