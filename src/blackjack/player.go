package blackjack

import "gaming/bankroll"

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
	Bankroll() bankroll.MoneyHolder
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
		HandHolder: NewHandHolder(),
		Better: NewBetter(bankroll.NewMoneyHolder(), bettingStrategy),
		playStrategy: playStrategy,
	}
}

func (player *playerImpl) PlayStrategy() PlayStrategy {
	return player.playStrategy
}
