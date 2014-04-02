/**
 * Date: 4/2/14
 * Time: 4:15 PM
 * @author jack 
 */
package blackjack

type BettingStrategy interface {
	GetUnitsToBet() float64
}

type consistentBettingStrategy struct {
	units float64
}

func NewConsistentBettingStrategy(units float64) BettingStrategy {
	return &consistentBettingStrategy{units}
}

func (this *consistentBettingStrategy) GetUnitsToBet() float64 {
	return this.units
}
