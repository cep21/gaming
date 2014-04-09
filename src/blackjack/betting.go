/**
 * Date: 4/2/14
 * Time: 4:15 PM
 * @author jack 
 */
package blackjack

import "gaming/bankroll"

// A strategy to pick how much money to bet
type BettingStrategy interface {
	GetMoneyToBet() bankroll.Money
}

type Better interface {
	Bankroll() bankroll.MoneyHolder
	BettingStrategy() BettingStrategy
}

type betterImpl struct {
	bankroll        bankroll.MoneyHolder
	bettingStrategy BettingStrategy
}

func (better *betterImpl) Bankroll() bankroll.MoneyHolder {
	return better.bankroll
}
func (better *betterImpl) BettingStrategy() BettingStrategy {
	return better.bettingStrategy
}

func NewBetter(bankroll bankroll.MoneyHolder, bettingStrategy BettingStrategy) Better {
	return &betterImpl{
		bankroll: bankroll,
		bettingStrategy: bettingStrategy,
	}
}

type consistentBettingStrategy struct {
	money bankroll.Money
}

func NewConsistentBettingStrategy(money bankroll.Money) BettingStrategy {
	return &consistentBettingStrategy{money: money}
}

func (this *consistentBettingStrategy) GetMoneyToBet() bankroll.Money {
	return this.money
}
