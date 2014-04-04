/**
 * Date: 4/4/14
 * Time: 10:55 AM
 * @author jack 
 */
package blackjack

type Bankroll interface {
	ChangeBankroll(amount float64)
	CurrentBankroll() float64
}

type bankrollImpl struct {
	currentBankroll float64
}

func (this *bankrollImpl) ChangeBankroll(amount float64) {
	this.currentBankroll += amount
}

func (this *bankrollImpl) CurrentBankroll() float64 {
	return this.currentBankroll
}

func NewBankroll(startingAmount float64) Bankroll {
	return &bankrollImpl{currentBankroll:startingAmount}
}
