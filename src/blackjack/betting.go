/**
 * Date: 4/2/14
 * Time: 4:15 PM
 * @author jack
 */
package blackjack

import "gaming"

// A strategy to pick how much money to bet
type BettingStrategy interface {
	GetMoneyToBet() gaming.Money
}

type Better interface {
	Bankroll() gaming.MoneyHolder
	BettingStrategy() BettingStrategy
}

func NewBetter(bankroll gaming.MoneyHolder, bettingStrategy BettingStrategy) Better {
	return &betterImpl{
		bankroll:        bankroll,
		bettingStrategy: bettingStrategy,
	}
}

func NewConsistentBettingStrategy(money gaming.Money) BettingStrategy {
	return &consistentBettingStrategy{money: money}
}

func (this *consistentBettingStrategy) GetMoneyToBet() gaming.Money {
	return this.money
}


/// ---- Private implementations

type consistentBettingStrategy struct {
	money gaming.Money
}

type betterImpl struct {
	bankroll        gaming.MoneyHolder
	bettingStrategy BettingStrategy
}

func (better *betterImpl) Bankroll() gaming.MoneyHolder {
	return better.bankroll
}
func (better *betterImpl) BettingStrategy() BettingStrategy {
	return better.bettingStrategy
}

