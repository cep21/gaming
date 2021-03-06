package blackjack

import (
	"gaming"
	"strconv"
)

type Hand interface {
	Score() uint
	IsSoft() bool
	CanSplit() bool
	Size() uint
	Push(c Card)
	IsBlackjack() bool
	IsSplitHand() bool
	FirstCard() Card
	Cards() []Card
	MoneyInThisHand() gaming.MoneyHolder
	SplitNumber() uint
	LastAction() GameAction
	SetLastAction(GameAction)
	SplitHand(bankrollToDrawFrom gaming.MoneyHolder) (Hand, Hand, error)

	Bust() bool
	Clone(bankrollToDrawFrom gaming.MoneyHolder) Hand
}

type HandHolder interface {
	Hands() []Hand
	SetHands(hands []Hand)
}

type handHolderImpl struct {
	hands []Hand
}

func NewHandHolder() HandHolder {
	return &handHolderImpl{hands: nil}
}

func (holder *handHolderImpl) Hands() []Hand {
	return holder.hands
}
func (holder *handHolderImpl) SetHands(hands []Hand) {
	holder.hands = hands
}

type handImpl struct {
	cards       []Card
	aceCount    uint
	score       uint
	splitNumber uint
	money       gaming.MoneyHolder
	lastAction  GameAction
}

func (this *handImpl) Cards() []Card {
	return this.cards
}

func (this *handImpl) IsSplitHand() bool {
	return this.splitNumber > 0
}

func (this *handImpl) MoneyInThisHand() gaming.MoneyHolder {
	return this.money
}

func (this *handImpl) SetLastAction(lastAction GameAction) {
	this.lastAction = lastAction
}

func (this *handImpl) LastAction() GameAction {
	return this.lastAction
}

func (this *handImpl) SplitHand(bankrollToDrawFrom gaming.MoneyHolder) (Hand, Hand, error) {
	h1 := NewSplitHand(this.SplitNumber()+1, this.cards[0])
	bankrollToDrawFrom.TransferMoneyTo(h1.MoneyInThisHand(), this.MoneyInThisHand().CurrentBankroll())
	h2 := NewSplitHand(this.SplitNumber()+1, this.cards[1])
	this.MoneyInThisHand().TransferMoneyTo(h2.MoneyInThisHand(), this.MoneyInThisHand().CurrentBankroll())
	return h1, h2, nil
}

func (this *handImpl) SplitNumber() uint {
	return this.splitNumber
}

func (this *handImpl) FirstCard() Card {
	if len(this.cards) == 0 {
		panic("Please check the first card first")
		//return nil
	} else {
		return this.cards[0]
	}
}

func (this *handImpl) Score() uint {
	if this.IsSoft() {
		return this.score + 10
	} else {
		return this.score
	}
}

func (this *handImpl) String() string {
	if this.IsBlackjack() {
		return "blackjack"
	}
	s := ""
	if this.IsSoft() {
		s += "soft "
	} else {
		s += "hard "
	}
	s += strconv.FormatUint(uint64(this.Score()), 10)
	return s
}

func (this *handImpl) Size() uint {
	return uint(len(this.cards))
}

func (this *handImpl) Clone(bankrollToDrawFrom gaming.MoneyHolder) Hand {
	newBankroll := gaming.NewMoneyHolder()
	bankrollToDrawFrom.TransferMoneyTo(newBankroll, this.MoneyInThisHand().CurrentBankroll())
	return &handImpl{
		cards:       this.cards,
		aceCount:    this.aceCount,
		score:       this.score,
		splitNumber: this.splitNumber,
		money:       newBankroll,
		lastAction:  this.lastAction,
	}
}

func (this *handImpl) IsBlackjack() bool {
	return this.Size() == 2 && this.Score() == 21
}

func (this *handImpl) CanSplit() bool {
	return this.Size() == 2 && this.cards[0].Score() == this.cards[1].Score()
}

func (this *handImpl) IsSoft() bool {
	return this.aceCount > 0 && this.score < 12
}

func (this *handImpl) Bust() bool {
	return this.Score() > 21
}

func (this *handImpl) Push(c Card) {
	this.cards = append(this.cards, c)
	this.score += c.Score()
	if c.Value() == Ace {
		this.aceCount++
	}
}

func NewHand(cards ...Card) Hand {
	return NewSplitHand(0, cards...)
}

func NewSplitHand(splitNumber uint, cards ...Card) Hand {
	h := &handImpl{
		money: gaming.NewMoneyHolder(),
		splitNumber: splitNumber,
	}
	for _, c := range cards {
		h.Push(c)
	}
	return h
}
