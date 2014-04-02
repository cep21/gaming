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

type randomPickDeck struct {
	storedCards []Card
	defaultSuit gaming.Suit
	r *rand.Rand
	countPerValue []uint
}

func NewRandomPickDeck(r *rand.Rand, defaultSuit gaming.Suit, number_of_decks uint) CardSet {
	countPerValue := []uint{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i:=0;i<len(countPerValue);i++ {
		countPerValue[i] = number_of_decks * 4
	}
	return &randomPickDeck{[]Card{}, defaultSuit, r, countPerValue}
}

func (this *randomPickDeck) Cards() []Card {
	panic("Don't look at the cards in a random pick deck!")
}

func (this *randomPickDeck) Peek() Card {
	if len(this.storedCards) == 0 {
		this.storedCards = append(this.storedCards, this.generateNextCard())
	}
	return this.storedCards[0]
}

func (this *randomPickDeck) generateNextCard() Card {
	sum := this.CardsLeft()
	if sum == 0 {
		return nil
	}
	selected_index := uint(this.r.Intn(int(sum)))
	for card_val := 0; card_val < len(this.countPerValue); card_val++ {
		if selected_index < this.countPerValue[card_val]  {
			this.countPerValue[card_val]--;
			return NewCard(this.defaultSuit, Values()[card_val])
		} else {
			selected_index -= this.countPerValue[card_val]
		}
	}
	panic("My logic is flawed.  This should never happen")
}

func (this *randomPickDeck) Pop() Card {
	if len(this.storedCards) == 0 {
		return this.generateNextCard()
	}
	c := this.storedCards[0]
	this.storedCards = this.storedCards[1:]
	return c
}

func (this *randomPickDeck) Push(c Card) {
	this.storedCards = append(this.storedCards, c)
}

func (this *randomPickDeck) CardsLeft() uint {
	count := uint(0)
	for i:=0;i<len(this.countPerValue);i++ {
		count += this.countPerValue[i]
	}
	return count
}

func (this *randomPickDeck) IsEmpty() bool {
	return this.CardsLeft() == 0
}

type CardsetFactory interface {
	CreateCardSet() CardSet
}

type infiniteCardSetFactory struct {
	r *rand.Rand
}

func (this *infiniteCardSetFactory) CreateCardSet() CardSet {
	return NewInfiniteDeck(this.r)
}

func NewInfiniteCardSetFactory(r *rand.Rand) CardsetFactory {
	return &infiniteCardSetFactory{r}
}

type clonedDeckStrategy struct {
	originalDeck CardSet
}

func (this *infiniteCardSetFactory) CreateCardSet() CardSet {
	return NewInfiniteDeck(this.r)
}

func NewClonedDeckStrategy(originalDeck CardSet) CardsetFactory {

}
