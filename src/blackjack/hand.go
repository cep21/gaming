package blackjack

type Hand interface {
CardSet
	Score() uint
	IsSoft() bool
	CanSplit() bool
	IsBlackjack() bool
	Bust() bool
}

type handImpl struct {
	cardSetImpl
	aceCount uint
	score uint
}

func (this *handImpl) Score() uint {
	if this.IsSoft() {
		return this.score + 10;
	} else {
		return this.score;
	}
}

func (this *handImpl) IsBlackjack() bool {
	return this.CardsLeft() == 2 && this.Score() == 21
}

func (this *handImpl) CanSplit() bool {
	return this.CardsLeft() == 2 && this.cards[0].Score() == this.cards[1].Score()
}

func (this *handImpl) IsSoft() bool {
	return this.aceCount > 0 && this.score < 12;
}

func (this *handImpl) Bust() bool {
	return this.Score() > 21
}

func (this *handImpl) Pop() Card {
	c := this.cardSetImpl.Pop()
	this.score -= c.Score()
	if c.Value() == Ace {
		this.aceCount--;
	}
	return c
}

func (this *handImpl) Push(c Card) {
	this.cardSetImpl.Push(c)
	this.score += c.Score()
	if c.Value() == Ace {
		this.aceCount++;
	}
}

func NewHand() Hand {
	return &handImpl{}
}

func NewHandWithCard(c Card) Hand {
	h := NewHand()
	h.Push(c)
	return h
}
