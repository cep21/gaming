package blackjack

import "gaming"

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
