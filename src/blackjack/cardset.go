package blackjack

import (
	"gaming"
	"math"
	"math/rand"
)

type CardSet interface {
	Cards() []Card
	Peek() Card
	Pop() Card
	Push(Card)
	CardsLeft() uint
	IsEmpty() bool
}

type cardSetImpl struct {
	cards []Card
}

func (this *cardSetImpl) Cards() []Card {
	return this.cards
}

func (this *cardSetImpl) Peek() Card {
	return this.cards[0]
}

func (this *cardSetImpl) Pop() Card {
	c := this.cards[0]
	this.cards = this.cards[1:]
	return c
}

func (this *cardSetImpl) Push(c Card) {
	this.cards = append(this.cards, c)
}

func (this *cardSetImpl) CardsLeft() uint {
	return uint(len(this.cards))
}

func (this *cardSetImpl) IsEmpty() bool {
	return this.CardsLeft() != 0
}

func NewDeck(cards []Card) CardSet {
	return &cardSetImpl{cards}
}

func SingleDeck() CardSet {
	return Decks(uint(1))
}

func Decks(num_decks uint) CardSet {
	cards := []Card{}
	for i := uint(0); i < uint(num_decks); i++ {
		for _, suit := range gaming.Suits() {
			for _, value := range Values() {
				cards = append(cards, NewCard(suit, value))
			}
		}
	}

	return NewDeck(cards)
}

type infiniteDeck struct {
	storedCards []Card
	r *rand.Rand
}

func (this *infiniteDeck) Cards() []Card {
	panic("Don't look at the cards in an infinite deck!")
}

func (this *infiniteDeck) Peek() Card {
	if len(this.storedCards) == 0 {
		this.generateNextCard()
	}
	return this.storedCards[0]
}

func (this *infiniteDeck) generateNextCard() {
	this.storedCards = append(this.storedCards, NewRandomCard(this.r))
}

func (this *infiniteDeck) Pop() Card {
	if len(this.storedCards) == 0 {
		this.generateNextCard()
	}
	c := this.storedCards[0]
	this.storedCards = this.storedCards[1:]
	return c
}

func (this *infiniteDeck) Push(c Card) {
	this.storedCards = append(this.storedCards, c)
}

func (this *infiniteDeck) CardsLeft() uint {
	return math.MaxUint64
}

func (this *infiniteDeck) IsEmpty() bool {
	return false;
}

func NewInfiniteDeck(rand *rand.Rand) CardSet {
	return &infiniteDeck{storedCards: []Card{}, r:rand}
}
