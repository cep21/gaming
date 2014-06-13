package blackjack

import "gaming"

type Dealer interface {
	HandHolder
	BankrolledStrategy
	HandDealer() HandDealer
}

type dealerImpl struct {
	HandHolder
	bankroll     gaming.MoneyHolder
	playStrategy PlayStrategy
	handDealer   HandDealer
}

func NewDealer(playStrategy PlayStrategy, handDealer HandDealer) Dealer {
	return &dealerImpl{
		HandHolder:   NewHandHolder(),
		bankroll:     gaming.NewMoneyHolder(),
		playStrategy: playStrategy,
		handDealer:   handDealer,
	}
}

func (dealer *dealerImpl) Bankroll() gaming.MoneyHolder {
	return dealer.bankroll
}
func (dealer *dealerImpl) PlayStrategy() PlayStrategy {
	return dealer.playStrategy
}
func (dealer *dealerImpl) HandDealer() HandDealer {
	return dealer.handDealer
}
