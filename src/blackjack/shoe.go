/**
 * Date: 4/3/14
 * Time: 12:52 PM
 * @author jack 
 */
package blackjack

import (
	"gaming"
	"math/rand"
	"math"
)

type Shoe interface {
	Pop() Card
	CardsLeft() uint
	Clone() Shoe
	Shuffle(r *rand.Rand) Shoe
}

type setShoeImpl struct {
	cards []Card
}

func (this *setShoeImpl) Pop() Card {
	c := this.cards[0]
	this.cards = this.cards[1:]
	return c
}

func (this *setShoeImpl) CardsLeft() uint {
	return uint(len(this.cards))
}

func (this *setShoeImpl) Clone() Shoe {
	v := []Card{}
	v = append(v, this.cards...)
	return &setShoeImpl{v}
}

func (this *setShoeImpl) Shuffle(r *rand.Rand) Shoe {
	for i := range this.cards {
		j := r.Intn(i + 1)
		this.cards[i], this.cards[j] = this.cards[j], this.cards[i]
	}
	return this
}

func NewShoe(cards []Card) Shoe {
	return &setShoeImpl{cards}
}

func Decks(num_decks uint) Shoe {
	cards := []Card{}
	for i := uint(0); i < uint(num_decks); i++ {
		for _, suit := range gaming.Suits() {
			for _, value := range Values() {
				cards = append(cards, NewCard(suit, value))
			}
		}
	}

	return NewShoe(cards)
}

type infiniteShoe struct {
	r *rand.Rand
}

func (this *infiniteShoe) Pop() Card {
	return NewRandomCard(this.r)
}

func (this *infiniteShoe) CardsLeft() uint {
	return math.MaxUint64
}

func (this *infiniteShoe) Clone() Shoe {
	return this
}

func (this *infiniteShoe) Shuffle(r *rand.Rand) Shoe {
	return this
}

func NewInfiniteShoe(rand *rand.Rand) Shoe {
	return &infiniteShoe{r:rand}
}

type randomPickShoe struct {
	defaultSuit gaming.Suit
	r *rand.Rand
	countPerValue []uint
}

func NewRandomPickShoe(r *rand.Rand, defaultSuit gaming.Suit, number_of_decks uint) Shoe {
	countPerValue := []uint{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i:=0;i<len(countPerValue);i++ {
		countPerValue[i] = number_of_decks * 4
	}
	return &randomPickShoe{defaultSuit, r, countPerValue}
}

func (this *randomPickShoe) Pop() Card {
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

func (this *randomPickShoe) CardsLeft() uint {
	count := uint(0)
	for i:=0;i<len(this.countPerValue);i++ {
		count += this.countPerValue[i]
	}
	return count
}

func (this *randomPickShoe) Clone() Shoe {
	countPerValue := []uint{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i:=0;i<len(countPerValue);i++ {
		countPerValue[i] = this.countPerValue[i]
	}
	return &randomPickShoe{this.defaultSuit, this.r, countPerValue}
}

func (this *randomPickShoe) Shuffle(r *rand.Rand) Shoe {
	return this
}


type ShoeFactory interface {
	CreateShoe() Shoe
}

type clonedShoeFactory struct {
	originalShoe Shoe
	r *rand.Rand
}

func (this *clonedShoeFactory) CreateShoe() Shoe {
	return this.originalShoe.Clone().Shuffle(this.r)
}

func NewClonedDeckFactory(originalShoe Shoe, r *rand.Rand) ShoeFactory {
	return &clonedShoeFactory{originalShoe, r}
}
