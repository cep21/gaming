package blackjack

import "gaming"

type Player interface {
	HandHolder
	Better
	PlayStrategy() PlayStrategy
}

type HandHolderBetter interface {
	HandHolder
	Better
}

type BankrolledStrategy interface {
	Bankroll() gaming.MoneyHolder
	PlayStrategy() PlayStrategy
}

type playerImpl struct {
	HandHolder
	Better
	playStrategy PlayStrategy
	hands        []Hand
}

func NewPlayer(bettingStrategy BettingStrategy, playStrategy PlayStrategy) Player {
	return &playerImpl{
		HandHolder:   NewHandHolder(),
		Better:       NewBetter(gaming.NewMoneyHolder(), bettingStrategy),
		playStrategy: playStrategy,
	}
}

func (player *playerImpl) PlayStrategy() PlayStrategy {
	return player.playStrategy
}
