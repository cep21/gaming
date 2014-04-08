package blackjack

import (
	"strconv"
	"gaming/bankroll"
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
	MoneyInThisHand() bankroll.MoneyHolder
	SplitNumber() uint
	LastAction() GameAction
	SetLastAction(GameAction)
	SplitHand(bankrollToDrawFrom bankroll.MoneyHolder) (Hand, Hand, error)

	Bust() bool
	Clone(bankrollToDrawFrom bankroll.MoneyHolder) Hand
}

type handImpl struct {
	cards       []Card
	aceCount    uint
	score       uint
	splitNumber uint
	money       bankroll.MoneyHolder
	lastAction  GameAction
}

func (this *handImpl) Cards() []Card {
	return this.cards
}

func (this *handImpl) IsSplitHand() bool {
	return this.splitNumber > 0
}

func (this *handImpl) MoneyInThisHand() bankroll.MoneyHolder {
	return this.money
}

func (this *handImpl) SetLastAction(lastAction GameAction) {
	this.lastAction = lastAction
}

func (this *handImpl) LastAction() GameAction {
	return this.lastAction
}

func (this *handImpl) SplitHand(bankrollToDrawFrom bankroll.MoneyHolder) (Hand, Hand, error) {
	newBankroll := bankroll.NewMoneyHolder()
	bankrollToDrawFrom.TransferMoneyTo(newBankroll, this.MoneyInThisHand().CurrentBankroll())
	h1 := NewHand(this.cards[0])
	h1.MoneyInThisHand()
	h2 := NewHand(this.cards[1])
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
		return this.score + 10;
	} else {
		return this.score;
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

func (this *handImpl) Clone(bankrollToDrawFrom bankroll.MoneyHolder) Hand {
	newBankroll := bankroll.NewMoneyHolder()
	bankrollToDrawFrom.TransferMoneyTo(newBankroll, this.MoneyInThisHand().CurrentBankroll())
	return &handImpl{
		cards: this.cards,
		aceCount: this.aceCount,
		score: this.score,
		splitNumber: this.splitNumber,
		money: newBankroll,
		lastAction: this.lastAction,
	}
}

func (this *handImpl) IsBlackjack() bool {
	return this.Size() == 2 && this.Score() == 21
}

func (this *handImpl) CanSplit() bool {
	return this.Size() == 2 && this.cards[0].Score() == this.cards[1].Score()
}

func (this *handImpl) IsSoft() bool {
	return this.aceCount > 0 && this.score < 12;
}

func (this *handImpl) Bust() bool {
	return this.Score() > 21
}

func (this *handImpl) Push(c Card) {
	this.cards = append(this.cards, c)
	this.score += c.Score()
	if c.Value() == Ace {
		this.aceCount++;
	}
}

func NewHand(cards ...Card) Hand {
	h := &handImpl{
		money: bankroll.NewMoneyHolder(),
	}
	for _, c := range cards {
		h.Push(c)
	}
	return h
}
