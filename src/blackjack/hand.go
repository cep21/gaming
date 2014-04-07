package blackjack

import "strconv"

type Hand interface {
	Score() uint
	IsSoft() bool
	CanSplit() bool
	Size() uint
	Push(c Card)
	IsBlackjack() bool
//	IsSplitHand() bool
	FirstCard() Card
	Cards() []Card
	MoneyInThisHand() MoneyHolder
	SplitNumber() uint
	LastAction() GameAction
	SetLastAction(GameAction)
	SplitHand(bankrollToDrawFrom MoneyHolder) (Hand, Hand, error)

	Bust() bool
	Clone(bankrollToDrawFrom MoneyHolder) Hand
}

type handImpl struct {
	cards []Card
	aceCount uint
	score uint
	splitNumber uint
	money MoneyHolder
	lastAction GameAction
}

func (this *handImpl) Cards() []Card {
	return this.cards
}

func (this *handImpl) MoneyInThisHand() MoneyHolder {
	return this.money
}

func (this *handImpl) SetLastAction(lastAction GameAction) {
	this.lastAction = lastAction
}

func (this *handImpl) LastAction() GameAction {
	return this.lastAction
}

func (this *handImpl) SplitHand(bankrollToDrawFrom MoneyHolder) (Hand, Hand, error) {

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

func (this *handImpl) Clone(bankrollToDrawFrom MoneyHolder) Hand {
	newBankroll := NewMoneyHolder()
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
	h := &handImpl{}
	for _, c := range cards {
		h.Push(c)
	}
	return h
}
