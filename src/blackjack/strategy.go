/**
 * Date: 4/1/14
 * Time: 3:46 PM
 * @author jack 
 */
package blackjack

type ShouldHitStrategy interface {
	ShouldHit(currentHand Hand, opponentHand Hand) bool
}

type dealerHitStrategy struct {
	hitOnSoft17 bool
}

func (this* dealerHitStrategy) ShouldHit(currentHand Hand, opponentHand Hand) bool {
	score := currentHand.Score()
	if score < 17 {
		return true
	} else if score > 17 {
		return false;
	} else {
		return this.hitOnSoft17
	}
}

func NewDealerStrategy(hitOnSoft17 bool) ShouldHitStrategy {
	return &dealerHitStrategy{hitOnSoft17: hitOnSoft17}
}
