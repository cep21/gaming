/**
 * Date: 4/1/14
 * Time: 3:46 PM
 * @author jack 
 */
package blackjack

type ShouldHitStrategy interface {
	ShouldHit(currentHand Hand, shownCard Card) bool
}

type dealerHitStrategy struct {
	hitOnSoft17 bool
}

func (this* dealerHitStrategy) ShouldHit(currentHand Hand, shownCard Card) bool {
	score := currentHand.Score()
	if score == 17 {
		return this.hitOnSoft17 && currentHand.IsSoft()
	} else {
		return score < 17;
	}
}

func NewDealerStrategy(hitOnSoft17 bool) ShouldHitStrategy {
	return &dealerHitStrategy{hitOnSoft17: hitOnSoft17}
}

// Always hit
type alwaysHitStrategy struct {
}

func (this* alwaysHitStrategy) ShouldHit(currentHand Hand, shownCard Card) bool {
	return true;
}

func NewAlwaysHitStrategy() ShouldHitStrategy {
	return &alwaysHitStrategy{}
}

// Always stand
type alwaysStandStrategy struct {
}

func (this* alwaysStandStrategy) ShouldHit(currentHand Hand, shownCard Card) bool {
	return false;
}

func NewAlwaysStandStrategy() ShouldHitStrategy {
	return &alwaysStandStrategy{}
}

// Never bust
type neverBustStrategy struct {
	shouldHitSoft bool
}

func (this* neverBustStrategy) ShouldHit(currentHand Hand, shownCard Card) bool {
	return (currentHand.IsSoft() && this.shouldHitSoft) || currentHand.Score() < 12;
}

func NewNeverBustStrategy(should_hit_soft bool) ShouldHitStrategy {
	return &neverBustStrategy{should_hit_soft}
}

func PlayHandOnStrategy(currentHand Hand, shownCard Card, strategy ShouldHitStrategy, deck CardSet) {
	for ; !deck.IsEmpty(); {
		if currentHand.Bust() || !strategy.ShouldHit(currentHand, shownCard) {
			return;
		} else {
			c := deck.Pop()
			currentHand.Push(c)
		}
	}
}

// Hit on a hard score strategy
type hitOnAScoreStrategy struct {
	hardScoreToHit uint
}

func (this* hitOnAScoreStrategy) ShouldHit(currentHand Hand, shownCard Card) bool {
	return currentHand.IsSoft() || currentHand.Score() == this.hardScoreToHit
}

func NewHitOnAScoreStrategy(score_to_hit uint) ShouldHitStrategy {
	return &hitOnAScoreStrategy{score_to_hit}
}

