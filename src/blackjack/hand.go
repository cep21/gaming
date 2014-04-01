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
	score    uint
	aceCount uint
}

func (this *handImpl) Score() uint {
	acecount, score := this.aceCount, this.score
	for ; acecount > 0 && score > 21; {
		acecount--
		score -= 10
	}
	return score
}

func (this *handImpl) IsBlackjack() bool {
	return this.CardsLeft() == 2 && this.Score() == 21
}

func (this *handImpl) CanSplit() bool {
	return this.CardsLeft() == 2 && this.cards[0].Score() == this.cards[1].Score()
}

func (this *handImpl) IsSoft() bool {
	return this.Score() + this.aceCount*10 != this.score
}

func (this *handImpl) Bust() bool {
	return this.Score() > 21
}

func (this *handImpl) Pop() Card {
	c := this.cardSetImpl.Pop()
	this.score -= c.Score()
	if c.Value() == Ace {
		this.aceCount--
	}
	return c
}

func (this *handImpl) Push(c Card) {
	this.cardSetImpl.Push(c)
	this.score += c.Score()
	if c.Value() == Ace {
		this.aceCount++
	}
}

func NewHand() Hand {
	return &handImpl{}
}
