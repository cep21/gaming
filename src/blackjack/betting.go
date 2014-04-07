/**
 * Date: 4/2/14
 * Time: 4:15 PM
 * @author jack 
 */
package blackjack

// A strategy to pick how much money to bet
type BettingStrategy interface {
	GetMoneyToBet() Money
}

type consistentBettingStrategy struct {
	money Money
}

func NewConsistentBettingStrategy(money Money) BettingStrategy {
	return &consistentBettingStrategy{money: money}
}

func (this *consistentBettingStrategy) GetMoneyToBet() Money {
	return this.money
}
