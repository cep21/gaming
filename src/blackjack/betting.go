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

type consistentBettingStrategy struct {
	money bankroll.Money
}

func NewConsistentBettingStrategy(money bankroll.Money) BettingStrategy {
	return &consistentBettingStrategy{money: money}
}

func (this *consistentBettingStrategy) GetMoneyToBet() bankroll.Money {
	return this.money
}
