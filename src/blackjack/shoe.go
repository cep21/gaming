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
	"errors"
)

var ERR_CARD_NOT_IN_SHOE = errors.New("Unable to find card in the shoe")
var ERR_SHOE_EMPTY = errors.New("Tried to take a card from an empty shoe")

type Shoe interface {
	Pop() (Card, error)
	CardsLeft() uint
	TakeValueFromShoe(valueToTake Value) (Card, error)
	StartingCardCount() uint
	Clone() Shoe
	Shuffle(r *rand.Rand) Shoe
}

func ShoePenetration(s Shoe) float64 {
	return 1 - float64(s.CardsLeft()) / float64(s.StartingCardCount())
}

type setShoeImpl struct {
	cards []Card
	startingCardCount uint
}

func (this *setShoeImpl) Pop() (Card, error) {
	if len(this.cards) == 0 {
		return nil, ERR_SHOE_EMPTY
	}
	c := this.cards[0]
	this.cards = this.cards[1:]
	return c, nil
}

func (this *setShoeImpl) CardsLeft() uint {
	return uint(len(this.cards))
}

func (this *setShoeImpl) StartingCardCount() uint {
	return this.startingCardCount
}

func (this *setShoeImpl) Clone() Shoe {
	v := []Card{}
	v = append(v, this.cards...)
	return &setShoeImpl{v, this.startingCardCount}
}

func (this *setShoeImpl) Shuffle(r *rand.Rand) Shoe {
	for i := range this.cards {
		j := r.Intn(i + 1)
		this.cards[i], this.cards[j] = this.cards[j], this.cards[i]
	}
	return this
}

func (this *setShoeImpl) TakeValueFromShoe(valueToTake Value) (Card, error) {
	for i, c:= range this.cards {
		if c.Value() == valueToTake {
			ret := c
			this.cards = append(this.cards[:i], this.cards[i+1:]...)
			return ret, nil
		}
	}
	return nil, ERR_CARD_NOT_IN_SHOE
}

func NewShoe(cards ...Card) Shoe {
	return &setShoeImpl{cards, uint(len(cards))}
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

	return NewShoe(cards...)
}

type infiniteShoe struct {
	r *rand.Rand
}

func (this *infiniteShoe) Pop() (Card, error) {
	return NewRandomCard(this.r), nil
}

func (this *infiniteShoe) CardsLeft() uint {
	return math.MaxUint64
}

func (this *infiniteShoe) StartingCardCount() uint {
	return math.MaxUint64
}

func (this *infiniteShoe) Clone() Shoe {
	return this
}

func (this *infiniteShoe) Shuffle(r *rand.Rand) Shoe {
	return this
}


func (this *infiniteShoe) TakeValueFromShoe(valueToTake Value) (Card, error) {
	return NewCard(gaming.RandomSuit(this.r), valueToTake), nil
}

func NewInfiniteShoe(rand *rand.Rand) Shoe {
	return &infiniteShoe{r:rand}
}

type randomPickShoe struct {
	r *rand.Rand
	countPerValue []uint
	suitsPerValue [][]gaming.Suit
	startingSize uint
	currentSize uint
}

func NewRandomPickShoe(r *rand.Rand, number_of_decks uint) Shoe {
	countPerValue := []uint{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	suitsPerValue := make([][]gaming.Suit, 13)
	all_suits := gaming.Suits()
	for i:=0;i<len(countPerValue);i++ {
		countPerValue[i] = number_of_decks * 4
		for j:=uint(0);j<4*number_of_decks;j++ {
			suitsPerValue[i] = append(suitsPerValue[i], all_suits[j%4])
		}
	}
	return &randomPickShoe{
		r: r,
		countPerValue: countPerValue,
		suitsPerValue: suitsPerValue,
		startingSize: 4 * number_of_decks * 13,
		currentSize: 4 * number_of_decks * 13,
	}
}

func (this *randomPickShoe) TakeValueFromShoe(valueToTake Value) (Card, error) {
	if this.countPerValue[valueToTake.Index()] == 0 {
		return nil, ERR_CARD_NOT_IN_SHOE
	}
	suit := this.suitsPerValue[valueToTake.Index()][this.countPerValue[valueToTake.Index()] - 1]
	this.countPerValue[valueToTake.Index()]--
	this.currentSize--
	return NewCard(suit, valueToTake), nil
}

func (this *randomPickShoe) Pop() (Card, error) {
	old_number_of_cards_left := this.currentSize
	if old_number_of_cards_left == 0 {
		return nil, ERR_SHOE_EMPTY
	}

	selected_index := uint(this.r.Intn(int(old_number_of_cards_left)))
	for card_val := 0; card_val < len(this.countPerValue); card_val++ {
		if selected_index < this.countPerValue[card_val]  {
			return this.TakeValueFromShoe(Values()[card_val])
		} else {
			selected_index -= this.countPerValue[card_val]
		}
	}
	panic("My logic is flawed.  This should never happen")
}

func (this *randomPickShoe) CardsLeft() uint {
	return this.currentSize
}

func (this *randomPickShoe) StartingCardCount() uint {
	return this.startingSize
}

func (this *randomPickShoe) Clone() Shoe {
	countPerValue := []uint{}
	countPerValue = append(countPerValue, this.countPerValue...)
	suitsPerValue := make([][]gaming.Suit, 13)
	for i:=0;i<len(countPerValue);i++ {
		suitsPerValue[i] = append(suitsPerValue[i], this.suitsPerValue[i]...)
	}
	return &randomPickShoe{
		r: this.r,
		countPerValue: countPerValue,
		suitsPerValue: suitsPerValue,
		startingSize: this.startingSize,
		currentSize: this.currentSize,
	}
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
