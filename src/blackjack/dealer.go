package blackjack

import "gaming/bankroll"

type Dealer interface {
	HandHolder
	BankrolledStrategy
	HandDealer() HandDealer
}

type dealerImpl struct {
	HandHolder
	bankroll     bankroll.MoneyHolder
	playStrategy PlayStrategy
	handDealer   HandDealer
}

func NewDealer(playStrategy PlayStrategy, handDealer HandDealer) Dealer {
	return &dealerImpl{
		HandHolder:   NewHandHolder(),
		bankroll:     bankroll.NewMoneyHolder(),
		playStrategy: playStrategy,
		handDealer:   handDealer,
	}
}

func (dealer *dealerImpl) Bankroll() bankroll.MoneyHolder {
	return dealer.bankroll
}
func (dealer *dealerImpl) PlayStrategy() PlayStrategy {
	return dealer.playStrategy
}
func (dealer *dealerImpl) HandDealer() HandDealer {
	return dealer.handDealer
}
