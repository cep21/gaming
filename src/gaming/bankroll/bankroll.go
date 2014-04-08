/**
 * Date: 4/4/14
 * Time: 10:55 AM
 * @author jack 
 */
package bankroll

// A currency that you can bet and win
type Money float64

// An object that holds some amount of money
type MoneyHolder interface {
	// Transfers money from this object that holds money to another object that holds money
	TransferMoneyTo(MoneyHolder, Money)
	// Gets the current amount of money in this holder
	CurrentBankroll() Money
	// A private interface to give a holder money
	giveMoney(Money)
}

type moneyHolderImpl struct {
	currentBankroll Money
}

func (this *moneyHolderImpl) TransferMoneyTo(moneyHolder MoneyHolder, amount Money) {
	this.currentBankroll -= amount
	moneyHolder.giveMoney(amount)
}

func (this *moneyHolderImpl) CurrentBankroll() Money {
	return this.currentBankroll
}

func (this *moneyHolderImpl) giveMoney(money Money) {
	this.currentBankroll += money
}

func NewMoneyHolder() MoneyHolder {
	return &moneyHolderImpl{currentBankroll:Money(0)}
}
